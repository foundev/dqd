# DQD Development Guidelines for Claude & AI Agents

This document contains coding standards, best practices, and architectural guidelines for AI agents (Claude Code, Cursor, etc.) working on the DQD (Dremio Query Doctor) project.

## Table of Contents

- [Go Coding Standards](#go-coding-standards)
- [Error Handling](#error-handling)
- [Logging](#logging)
- [File I/O Patterns](#file-io-patterns)
- [HTTP Handlers](#http-handlers)
- [Testing](#testing)
- [Project Architecture](#project-architecture)
- [Common Antipatterns to Avoid](#common-antipatterns-to-avoid)

---

## Go Coding Standards

### Use Structured Logging (slog)

**ALWAYS** use `log/slog` for logging, never `log` package.

❌ **WRONG:**
```go
import "log"

log.Printf("Processing file: %s", filename)
log.Fatal("Server failed")
```

✅ **CORRECT:**
```go
import "log/slog"

slog.Info("processing file", "filename", filename, "size", size)
slog.Error("server failed", "error", err)
```

**Benefits:**
- Structured logging for better parsing
- Key-value pairs for context
- Log levels (Debug, Info, Warn, Error)
- Easy integration with log aggregators

### Error Handling

#### Never Ignore Errors

❌ **WRONG:**
```go
file.Close()  // Ignores potential error
```

✅ **CORRECT:**
```go
if err := file.Close(); err != nil {
    slog.Error("failed to close file", "error", err)
}
```

#### Don't Defer Functions That Return Errors

Deferring functions that return errors can hide problems and prevent proper cleanup.

❌ **WRONG:**
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()  // Error is ignored!
```

✅ **CORRECT - Option 1 (Named Return):**
```go
func processFile(path string) (err error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            slog.Error("failed to close file", "path", path, "error", closeErr)
            if err == nil {
                err = closeErr
            }
        }
    }()

    // Process file...
    return nil
}
```

✅ **CORRECT - Option 2 (Explicit Close):**
```go
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }

    // Process file...

    if err := file.Close(); err != nil {
        slog.Error("failed to close file", "path", path, "error", err)
        return fmt.Errorf("closing file: %w", err)
    }

    return nil
}
```

#### Wrap Errors with Context

Use `fmt.Errorf` with `%w` to wrap errors:

```go
if err := processData(data); err != nil {
    return fmt.Errorf("processing profile data: %w", err)
}
```

---

## Logging

### Structured Logging Patterns

#### Log Levels

Use appropriate log levels:

```go
// Debug - verbose information for debugging
slog.Debug("parsing operator", "id", opID, "type", opType)

// Info - general informational messages
slog.Info("server started", "port", port, "version", version)

// Warn - warning messages for unexpected but recoverable situations
slog.Warn("deprecated feature used", "feature", name)

// Error - error messages for failures
slog.Error("failed to parse profile", "error", err, "file", filename)
```

#### Context in Logs

Include relevant context as key-value pairs:

```go
slog.Info("request processed",
    "method", r.Method,
    "path", r.URL.Path,
    "status", status,
    "duration", duration,
    "client_ip", r.RemoteAddr,
)
```

#### Error Logging

Always include the error and context:

```go
if err != nil {
    slog.Error("database query failed",
        "error", err,
        "query", queryName,
        "params", params,
    )
    return err
}
```

---

## File I/O Patterns

### Reading Files

❌ **WRONG:**
```go
file, _ := os.Open("data.json")  // Ignoring error
defer file.Close()               // Ignoring close error
data, _ := io.ReadAll(file)      // Ignoring error
```

✅ **CORRECT:**
```go
func readDataFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("opening file: %w", err)
    }

    data, err := io.ReadAll(file)
    if err != nil {
        // Close file before returning error
        if closeErr := file.Close(); closeErr != nil {
            slog.Error("failed to close file after read error",
                "path", path,
                "closeError", closeErr,
            )
        }
        return nil, fmt.Errorf("reading file: %w", err)
    }

    if err := file.Close(); err != nil {
        slog.Error("failed to close file", "path", path, "error", err)
        return nil, fmt.Errorf("closing file: %w", err)
    }

    return data, nil
}
```

### Multipart File Uploads

For HTTP multipart files, explicit close is preferred:

❌ **WRONG:**
```go
file, _, err := r.FormFile("upload")
if err != nil {
    return err
}
defer file.Close()  // Ignores error
```

✅ **CORRECT:**
```go
file, header, err := r.FormFile("upload")
if err != nil {
    return fmt.Errorf("getting form file: %w", err)
}

data, err := io.ReadAll(file)
if err != nil {
    if closeErr := file.Close(); closeErr != nil {
        slog.Error("failed to close uploaded file",
            "filename", header.Filename,
            "closeError", closeErr,
        )
    }
    return fmt.Errorf("reading upload: %w", err)
}

if err := file.Close(); err != nil {
    slog.Error("failed to close uploaded file",
        "filename", header.Filename,
        "error", err,
    )
    return fmt.Errorf("closing upload: %w", err)
}
```

---

## HTTP Handlers

### Handler Structure

```go
func HandleUpload(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    // Parse request
    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        slog.Error("failed to parse form", "error", err)
        WriteError(w, "Invalid upload", http.StatusBadRequest)
        return
    }

    // Process request
    result, err := processUpload(r)
    if err != nil {
        slog.Error("upload processing failed", "error", err)
        WriteError(w, "Processing failed", http.StatusInternalServerError)
        return
    }

    // Write response
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    if _, err := w.Write([]byte(result)); err != nil {
        slog.Error("failed to write response", "error", err)
        return
    }

    slog.Info("request completed",
        "method", r.Method,
        "path", r.URL.Path,
        "duration", time.Since(start),
    )
}
```

### Error Responses

Create consistent error response helpers:

```go
func WriteError(w http.ResponseWriter, message string, statusCode int) {
    slog.Warn("returning error response",
        "status", statusCode,
        "message", message,
    )

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    response := ErrorResponse{
        Error:   http.StatusText(statusCode),
        Message: message,
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        slog.Error("failed to encode error response", "error", err)
    }
}
```

---

## Testing

### Unit Tests

Write tests for all public functions:

```go
func TestDetectFileType(t *testing.T) {
    tests := []struct {
        name     string
        filename string
        want     FileType
    }{
        {"JSON file", "profile.json", FileTypeJSON},
        {"ZIP file", "data.zip", FileTypeZIP},
        {"TAR.GZ file", "archive.tar.gz", FileTypeTarGZ},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := DetectFileType(tt.filename)
            if got != tt.want {
                t.Errorf("DetectFileType(%q) = %v, want %v",
                    tt.filename, got, tt.want)
            }
        })
    }
}
```

### Table-Driven Tests

Use table-driven tests for multiple test cases:

```go
func TestParseProfile(t *testing.T) {
    testCases := []struct {
        name      string
        input     []byte
        wantError bool
    }{
        {
            name:      "valid profile",
            input:     loadTestData("valid_profile.json"),
            wantError: false,
        },
        {
            name:      "invalid JSON",
            input:     []byte("{invalid}"),
            wantError: true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            _, err := ParseProfile(tc.input)
            if (err != nil) != tc.wantError {
                t.Errorf("ParseProfile() error = %v, wantError %v",
                    err, tc.wantError)
            }
        })
    }
}
```

---

## Project Architecture

### Package Structure

```
backend-go/
├── cmd/
│   └── server/          # Main application entry points
│       └── main.go
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data structures
│   ├── services/        # Business logic
│   └── fileutil/        # File utilities
├── pkg/                 # Public libraries (if any)
└── go.mod
```

### Dependency Injection

Prefer dependency injection over global state:

❌ **WRONG:**
```go
var db *sql.DB  // Global variable

func GetUser(id int) (*User, error) {
    return db.QueryRow("SELECT...", id)
}
```

✅ **CORRECT:**
```go
type UserService struct {
    db *sql.DB
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.db.QueryRow("SELECT...", id)
}
```

### Handler Dependencies

```go
type ProfileHandler struct {
    logger  *slog.Logger
    service *ProfileService
}

func NewProfileHandler(logger *slog.Logger, service *ProfileService) *ProfileHandler {
    return &ProfileHandler{
        logger:  logger,
        service: service,
    }
}

func (h *ProfileHandler) Handle(w http.ResponseWriter, r *http.Request) {
    h.logger.Info("handling profile request", "path", r.URL.Path)
    // ...
}
```

---

## Common Antipatterns to Avoid

### 1. Ignoring Errors

❌ **NEVER DO THIS:**
```go
file.Close()
json.Unmarshal(data, &result)
w.Write([]byte("response"))
```

### 2. Bare Returns

❌ **AVOID:**
```go
func process() (result string, err error) {
    result = "done"
    return  // Unclear what's being returned
}
```

✅ **PREFER:**
```go
func process() (string, error) {
    result := "done"
    return result, nil  // Explicit
}
```

### 3. Panic in Libraries

❌ **DON'T PANIC:**
```go
func parseConfig(path string) Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err)  // Bad!
    }
    // ...
}
```

✅ **RETURN ERRORS:**
```go
func parseConfig(path string) (Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return Config{}, fmt.Errorf("reading config: %w", err)
    }
    // ...
}
```

### 4. Silent Failures

❌ **WRONG:**
```go
if err := save(data); err != nil {
    // Nothing - error is lost!
}
```

✅ **CORRECT:**
```go
if err := save(data); err != nil {
    slog.Error("failed to save data", "error", err)
    return fmt.Errorf("saving data: %w", err)
}
```

### 5. Context Ignoring

❌ **WRONG:**
```go
func slowOperation() error {
    time.Sleep(10 * time.Second)  // No cancellation possible
    return nil
}
```

✅ **CORRECT:**
```go
func slowOperation(ctx context.Context) error {
    select {
    case <-time.After(10 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

---

## Migration-Specific Guidelines

### Java to Go Mapping

When porting Java code:

1. **Getters/Setters** → Public fields or methods
2. **Exceptions** → Multiple return values with error
3. **null** → Pointer types or zero values
4. **ArrayList** → Slices
5. **HashMap** → Maps
6. **synchronized** → Mutexes or channels

### Logging Migration

Java (SLF4J):
```java
logger.info("Processing file: {}", filename);
logger.error("Failed to process", exception);
```

Go (slog):
```go
slog.Info("processing file", "filename", filename)
slog.Error("failed to process", "error", err)
```

### Error Handling Migration

Java:
```java
try {
    processFile(filename);
} catch (IOException e) {
    logger.error("Failed", e);
    throw new RuntimeException(e);
}
```

Go:
```go
if err := processFile(filename); err != nil {
    slog.Error("failed to process file", "error", err, "filename", filename)
    return fmt.Errorf("processing file %s: %w", filename, err)
}
```

---

## Code Review Checklist

Before submitting code, verify:

- [ ] All errors are handled (no `_` for errors)
- [ ] `slog` is used for all logging (not `log`)
- [ ] No deferred functions with ignored errors
- [ ] File closes are explicitly handled
- [ ] All public functions have tests
- [ ] Error messages include context
- [ ] No panics in library code
- [ ] Proper use of contexts for cancellation
- [ ] Structured logging with key-value pairs
- [ ] Error wrapping with `%w`

---

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [slog Package Documentation](https://pkg.go.dev/log/slog)
- [Error Handling in Go](https://go.dev/blog/error-handling-and-go)

---

## Questions?

If you're unsure about a pattern:
1. Check this guide first
2. Look at similar code in the codebase
3. Ask in PR review
4. Default to idiomatic Go patterns

**When in doubt, prefer explicitness over cleverness.**
