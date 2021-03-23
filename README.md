# new-collector-template

Boilerplate for creating new collectors

1. Hit the "Use this template" button to create a new repository from the template
1. Update the go mod file with the new root module: `go mod edit -module "github.com/rode/$MY_COLLECTOR"`
1. Find and replace any usages of "new-collector-template"
1. Pick an unallocated set of ports and update the `config` package to replace the defaults for `--http-port` and `--grpc-port`
