# Futil - File Utility
- Futil is a command-line application written in Go, providing useful tools for working with files.
## Project structure
The application is designed with the following structure:
```
futil/
├── cmd/
│   └── futil/
│       └── main.go
├── internal/
│   ├── linecount/
│   │   └── linecount.go
│   ├── checksum/
│   │   └── checksum.go
│   └── utils/
│       └── utils.go
├── go.mod
├── go.sum
├── myfile.txt
└── README.md
```
- `cmd/futil/main.go`: Main entry point of the application, using Cobra to handle the CLI.

- `internal/linecount/`: Contains the logic to count the number of lines in the file.

- `internal/checksum/`: Contains the logic to calculate the checksum of the file.

- `internal/utils/`: Contains utility functions used throughout the application.

## Third-party libraries
- [github.com/spf13/cobra](https://github.com/spf13/cobra): Used to build command-line interfaces.

## How to build and run the project

1. Clone repository:
   ```
   git clone https://github.com/your-username/futil.git
   cd futil
   ```

2. Build project:
   ```
   go build -o futil ./cmd/futil
   ```

3. Run project:
   ```
   ./futil [command] [flags]
   ```

   Example:
   ```
   ./futil linecount -f myfile.txt
   ./futil checksum -f myfile.txt --md5
   ```

## Features

- Count lines in file
- Calculate file checksum (supports md5, sha1, sha256)
- Read file from stdin
- Handle non-existent or invalid file cases
- Show version and help
- Show help for subcommands

## CI/CD

This project uses GitHub Actions for CI/CD. The workflow includes:
- Lint checks
- Running unit tests
- Building the application
- Creating automatic releases when new tags are created