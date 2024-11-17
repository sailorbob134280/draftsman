# Print this help message
default:
  just --list

# Install the pre-commit hooks
bootstrap:
  pre-commit install --hook-type commit-msg --hook-type pre-push

# Build the draftsman binary
build:
  mkdir -p build
  CGO_ENABLED=false go build -v -o build/draftsman main.go

# Run draftsman
run:
  go run main.go

# Run the tests
test:
  go test -v ./...

# Run the tests with coverage
coverage:
  mkdir -p build
  go test -coverprofile=build/coverage.out ./...
  go tool cover -html=build/coverage.out -o build/coverage.html

# Run the linter
lint:
  go vet ./...

# Clean the build directory
clean:
  rm -rf build

# Print this help message
help:
  just --list
