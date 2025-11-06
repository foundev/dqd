package fileutil

import (
	"testing"
)

func TestDetectFileType(t *testing.T) {
	tests := []struct {
		filename string
		want     FileType
	}{
		{"profile.json", FileTypeJSON},
		{"data.zip", FileTypeZIP},
		{"archive.tar.gz", FileTypeTarGZ},
		{"archive.tgz", FileTypeTarGZ},
		{"file.tar", FileTypeTar},
		{"file.gz", FileTypeGZ},
		{"unknown.txt", FileTypeUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got := DetectFileType(tt.filename)
			if got != tt.want {
				t.Errorf("DetectFileType(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
