# Copilot Instructions for json-go

## Project Overview

- This is a Go implementation of a JSON parser, ported from TypeScript. The main logic is in `internal/shared/`.
- The parser is built around token types (`ValueToken`, `ArrayToken`, `ObjectToken`, etc.) defined in `types.go`.
- Parsing functions for each JSON type are in separate files: `array.go`, `object.go`, `number.go`, `string.go`, etc.
- The parser uses a state machine approach for each type, with explicit modes (e.g., Scanning, Value, End).

## Key Files & Structure

- `internal/shared/types.go`: Defines all token and value types used throughout the parser.
- `internal/shared/json.go`: Entry point for parsing JSON strings (`Parse` function).
- `internal/shared/array.go`, `object.go`, `number.go`, `string.go`: Parsing logic for each JSON type.
- `internal/shared/helpers.go`: Utility functions for delimiter handling and regex matching.
- Tests for each type are in corresponding `*_test.go` files.

## Developer Workflows

- **Build:**
  - Build all binaries: `go build ./...`
- **Format:**
  - Format all code: `go fmt ./...`
- **Run CLI:**
  - Run the CLI: `go run cmd/cli/main.go`
- **Test:**
  - Run all tests: `go test ./...`
  - Individual test files follow Go's standard `*_test.go` convention.
- **Test Coverage:**
  - Measure coverage: `go test -cover ./...`
- **Debugging:**
  - Parsing errors use `panic` for unexpected input or state. Check error messages for mode/state context.

## Project-Specific Patterns

- Parsing functions use explicit state machines with `mode` variables and switch statements.
- Whitespace and delimiters are handled via regex (`regexp.MatchString`).
- Each parser returns a token struct with a `skip` field indicating how many characters were consumed.
- The parser expects well-formed JSON; error handling is strict and will panic on unexpected input.
- All parsing is done in-memory; no external dependencies except Go's standard library.

## Integration Points

- No external APIs or services; the parser is self-contained.
- The CLI directory (`cmd/cli/`) provides a command-line interface for parsing JSON input interactively.

## Example Usage

```go
import "json-go/internal/shared"

result := json.Parse(`{"key": [1, 2, 3]}`)
// result is a ValueToken
```

## Conventions

- All code is in the `json` package under `internal/shared`.
- Token structs always include a `skip` field for character advancement.
- Parsing functions are named `parse<Type>` (e.g., `parseArray`, `parseObject`).

---

If any section is unclear or missing important details, please provide feedback to improve these instructions.
