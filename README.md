# DirCleaner
Command line application to clean up files under directory to subdirectories by date.

## Usage
Program expects absolute(full) path to directory.
The `pwd` command is useful to get absolute path if you are located in directory which should be cleaned.

## Build binary
Run following command `go build -o bin/dir-cleaner` to build executable binary located in `bin` directory.
If missing execution permission run `chmod +x bin/dir-cleaner`.

## Run tests
Run following command: `go test -v ./test`