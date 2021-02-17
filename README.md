## Kheiron Test

Built on Ubuntu 20, tested on Ubuntu 20 and Windows 10. Should work the same on Mac, but I don't have a Mac so I haven't tested it.

### Prerequisites: 
- Go version 1.15 or higher
- Some program for compiling with Makefiles will improve experience, but this is not necessary
- The API does use one external go library (gorilla/mux), but go should download that for you when you try to build. 

### Compilation:
- If you can compile using Makefiles, simply run `make` in the top-level directory of the project. 
- If you wish to only build the prefix calculator, the infix calculator, or the API, use `make build-prefix`, `make build-infix` or `make build-api` respectively.
- If you cannot use the Makefile, simply run the go command invoked by the command in the Makefile you would like to run (minus the @ at the start). For example, to build the prefix calculator you can run `go build -o bin/prefix-calculator.exe cmd/prefix-calculator/main.go`

### Running:
- After building, the binaries live in the `bin` directory. Run them as you would any binary executable.

### Testing:
- Run by either using `make test` or `go test -v ./tests/...` in the top-level of the project.

### Pipeline
- This project has a Gitlab pipelie, because  
    1. I made it in Gitlab initally, and migrated it over.
    2. It takes about a minute to set up, and eliminates silly mistakes.