#!/bin/bash

set -euo pipefail

RED=$(tput setaf 1)
YELLOW=$(tput setaf 3)
GREEN=$(tput setaf 2)
DEFAULT=$(tput sgr0)

function to_snake_case() {
    echo "${1//-/_}"
}

function to_camel_case() {
    echo "$1" | awk -F - '{printf "%s", $1; for(i=2; i<=NF; i++) printf "%s", toupper(substr($i,1,1)) substr($i,2); print"";}'
}

function to_pascal_case() {
    camelCase=$(to_camel_case "$1")
    echo "${camelCase}" | awk '{print toupper(substr($0,0,1))substr($0,2)}'
}

function warn() {
  printf "%s\n" "${RED}$1${DEFAULT}"
}

function attention() {
  printf "%s\n" "${YELLOW}$1${DEFAULT}"
}

function info() {
  printf "%s\n" "${GREEN}$1${DEFAULT}"
}

function replace() {
  attention "Replacing instances of '$1' with '$2'"
  find .  \
    \( -name "*.go" -o -name "*.yml" -o -name "*.proto" -o -name "Dockerfile" -o -name "Makefile" -o -name "go.mod" \) \
    -exec sed -i '' 's/'"$1"'/'"$2"'/g' {} \;
}

templateProjectName="new-collector-template"
projectName="${1-}"

if [ -z "$projectName" ]; then
  warn "Must supply a hyphenated project name"
  warn "ex: ./template.sh my-collector-name"
  exit 1
fi

info "Updating template to use name '$projectName'"

warn "Removing generated protobuf code"
rm -f proto/v1alpha1/*.pb.go
rm -f proto/v1alpha1/*.pb.*.go

warn "Renaming protobuf file"
mv "proto/v1alpha1/${templateProjectName}.proto" "proto/v1alpha1/${projectName}.proto"

replace "$templateProjectName" "$projectName"

templateProjectNameCamelCase=$(to_camel_case "$templateProjectName")
projectNameCamelCase=$(to_camel_case "$projectName")
replace "$templateProjectNameCamelCase" "$projectNameCamelCase"

templateProjectNamePascalCase=$(to_pascal_case "$templateProjectName")
projectNamePascalCase=$(to_pascal_case "$projectName")
replace "$templateProjectNamePascalCase" "$projectNamePascalCase"

templateProjectNameSnakeCase=$(to_snake_case "$templateProjectName")
projectNameSnakeCase=$(to_snake_case "$projectName")
replace "$templateProjectNameSnakeCase" "$projectNameSnakeCase"

info "Regenerating protos"
make generate

info "Formatting"
make fmt

info "Building binary"
go build -v .

info "Running tests"
make test

info "Cleaning up binary"
rm "$projectName"

info "Updating README"
echo "# $projectName" > README.md

info "Build Docker image"
docker build -t "$projectName" .

info "Done, deleting this script"
rm -- "$0"
