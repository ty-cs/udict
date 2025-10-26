# GEMINI.md

## Project Overview

This project is a command-line interface (CLI) tool named `udict` written in Go. It allows users to fetch definitions from Urban Dictionary directly from their terminal. The tool is built using the `cobra` library for command management and `lipgloss` for styling the output, making it colorful and easy to read.

The application has a simple structure with a `main.go` entry point that executes the root command from the `cmd` package. The `cmd` package defines the CLI commands (`define`, `random`, `version`). The actual API interaction with Urban Dictionary is handled in the `internal/api` package, and the response formatting is done in the `internal/formatter` package.

## Building and Running

To build and run the project, you can use the standard Go commands.

### Build

```sh
go build
```

This will create an executable named `udict` in the project's root directory.

### Run

To run the application, you can use the following commands:

**Define a word:**

```sh
go run main.go define <word>
```

or with the built executable:

```sh
./udict define <word>
```

**Get a random word:**

```sh
go run main.go random
```

or with the built executable:

```sh
./udict random
```

### Flags

The `define` command supports a `--limit` (or `-l`) flag to specify the number of definitions to display (up to a maximum of 5).

```sh
go run main.go define <word> --limit 3
```

## Development Conventions

*   **Commands:** All CLI commands are defined in the `cmd` package, with each command in its own file (e.g., `define.go`, `random.go`).
*   **API Logic:** The logic for fetching data from the Urban Dictionary API is located in `internal/api/urban.go`.
*   **Formatting:** The output is styled using the `lipgloss` library. The formatting logic is in `internal/formatter/lipgloss.go`.
*   **Dependencies:** The project uses Go modules for dependency management. Dependencies are listed in the `go.mod` file.
