# new-collector-template

Boilerplate for creating new collectors.

## Usage

This template can be automatically reconfigured, given the name of the new collector:

1. Hit the "Use this template" button to create a new repository from the template.
1. Run `./template.sh collector-foo-bar`
1. Make any needed modifications under the `proto` directory for the collector's API
1. Pick an unallocated port and update the `config` package to replace the default for `--port`.
    - The `config` tests will need to be updated as well

## What's Included

- Example collector endpoint
- Flag configuration under `config`
- Boilerplate for creating a gRPC server & gRPC gateway in the `server` package and `main.go` 
- Dockerfile
- Unit tests, including a mock for the Rode client. 
- Protobuf code generation
- [GoReleaser](https://goreleaser.com/) config  
- Makefile targets for common tasks
    - `coverage`: Runs the tests and opens a browser window with the code coverage report
    - `fmt`: Runs `gofmt`
    - `fmtcheck`: Checks to see if `gofmt` needs to run
    - `generate`: Generates client and server code from the protobuf definitions    
    - `license`: Runs `addlicense` to populate license headers on source code
    - `mocks`: Generates a mock for the Rode client    
    - `test`: Runs unit tests
    - `vet`: Runs `go vet`
