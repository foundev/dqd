package fileutil

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// FileType represents the type of uploaded file
type FileType int

const (
	FileTypeUnknown FileType = iota
	FileTypeJSON
	FileTypeZIP
	FileTypeTarGZ
	FileTypeTar
	FileTypeGZ
)

// DetectFileType determines the file type from filename and content
func DetectFileType(filename string) FileType {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".json":
		return FileTypeJSON
	case ".zip":
		return FileTypeZIP
	case ".gz":
		// Check if it's .tar.gz
		if strings.HasSuffix(strings.ToLower(filename), ".tar.gz") {
			return FileTypeTarGZ
		}
		return FileTypeGZ
	case ".tgz":
		return FileTypeTarGZ
	case ".tar":
		return FileTypeTar
	default:
		return FileTypeUnknown
	}
}

// ExtractedFile represents a file extracted from an archive
type ExtractedFile struct {
	Name    string
	Content []byte
}

// ExtractFromZip extracts files from a ZIP archive
func ExtractFromZip(data []byte) ([]ExtractedFile, error) {
	reader, err := zip.NewReader(strings.NewReader(string(data)), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to open zip: %w", err)
	}

	var files []ExtractedFile
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file in zip: %w", err)
		}

		content, readErr := io.ReadAll(rc)

		// Always close and handle the error
		if closeErr := rc.Close(); closeErr != nil {
			slog.Error("failed to close file in zip",
				"filename", file.Name,
				"error", closeErr,
			)
			if readErr == nil {
				return nil, fmt.Errorf("closing file in zip: %w", closeErr)
			}
		}

		if readErr != nil {
			return nil, fmt.Errorf("failed to read file in zip: %w", readErr)
		}

		files = append(files, ExtractedFile{
			Name:    file.Name,
			Content: content,
		})
	}

	return files, nil
}

// ExtractFromTarGZ extracts files from a .tar.gz archive
func ExtractFromTarGZ(data []byte) (files []ExtractedFile, err error) {
	gzReader, err := gzip.NewReader(strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to open gzip: %w", err)
	}

	defer func() {
		if closeErr := gzReader.Close(); closeErr != nil {
			slog.Error("failed to close gzip reader", "error", closeErr)
			if err == nil {
				err = fmt.Errorf("closing gzip reader: %w", closeErr)
			}
		}
	}()

	files, err = extractFromTar(gzReader)
	return files, err
}

// ExtractFromTar extracts files from a .tar archive
func ExtractFromTar(data []byte) ([]ExtractedFile, error) {
	return extractFromTar(strings.NewReader(string(data)))
}

// extractFromTar is a helper that extracts from any tar reader
func extractFromTar(reader io.Reader) ([]ExtractedFile, error) {
	tarReader := tar.NewReader(reader)
	var files []ExtractedFile

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read tar: %w", err)
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		content, err := io.ReadAll(tarReader)
		if err != nil {
			return nil, fmt.Errorf("failed to read file in tar: %w", err)
		}

		files = append(files, ExtractedFile{
			Name:    header.Name,
			Content: content,
		})
	}

	return files, nil
}

// ExtractFromGZ extracts a single file from a .gz archive
func ExtractFromGZ(data []byte) (content []byte, err error) {
	gzReader, err := gzip.NewReader(strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to open gzip: %w", err)
	}

	defer func() {
		if closeErr := gzReader.Close(); closeErr != nil {
			slog.Error("failed to close gzip reader", "error", closeErr)
			if err == nil {
				err = fmt.Errorf("closing gzip reader: %w", closeErr)
			}
		}
	}()

	content, err = io.ReadAll(gzReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read gzip content: %w", err)
	}

	return content, nil
}

// ProcessUploadedFile handles an uploaded file and returns its content
// If it's an archive, it extracts and returns the first profile.json found
func ProcessUploadedFile(fileHeader *multipart.FileHeader) (data []byte, filename string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", fmt.Errorf("opening uploaded file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			slog.Error("failed to close uploaded file",
				"filename", fileHeader.Filename,
				"error", closeErr,
			)
			if err == nil {
				err = fmt.Errorf("closing uploaded file: %w", closeErr)
			}
		}
	}()

	data, err = io.ReadAll(file)
	if err != nil {
		return nil, "", fmt.Errorf("reading uploaded file: %w", err)
	}

	fileType := DetectFileType(fileHeader.Filename)

	switch fileType {
	case FileTypeJSON:
		return data, fileHeader.Filename, nil

	case FileTypeZIP:
		return extractProfileFromArchive(data, ExtractFromZip)

	case FileTypeTarGZ:
		return extractProfileFromArchive(data, ExtractFromTarGZ)

	case FileTypeTar:
		return extractProfileFromArchive(data, ExtractFromTar)

	case FileTypeGZ:
		content, err := ExtractFromGZ(data)
		if err != nil {
			return nil, "", err
		}
		return content, strings.TrimSuffix(fileHeader.Filename, ".gz"), nil

	default:
		// Assume it's JSON if we can't determine
		return data, fileHeader.Filename, nil
	}
}

type extractFunc func([]byte) ([]ExtractedFile, error)

func extractProfileFromArchive(data []byte, extract extractFunc) ([]byte, string, error) {
	files, err := extract(data)
	if err != nil {
		return nil, "", err
	}

	// Look for profile.json or any .json file
	for _, file := range files {
		if strings.HasSuffix(strings.ToLower(file.Name), "profile.json") {
			return file.Content, file.Name, nil
		}
	}

	// If no profile.json, return first .json file
	for _, file := range files {
		if strings.HasSuffix(strings.ToLower(file.Name), ".json") {
			return file.Content, file.Name, nil
		}
	}

	if len(files) > 0 {
		return files[0].Content, files[0].Name, nil
	}

	return nil, "", fmt.Errorf("no files found in archive")
}
