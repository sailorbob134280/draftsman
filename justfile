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
  go test -v -coverprofile=build/coverage.out ./...
  go tool cover -func=build/coverage.out
  go tool cover -html=build/coverage.out -o build/coverage.html

# Run all code checks
check: check-format vet

# Format all code
format:
  gofmt -w .

# Check the format of the code
check-format:
  if [ "$(gofmt -l . | wc -l)" -gt 0 ]; then exit 1; fi

# Run the vet tool
vet:
  go vet ./...

# Clean the build directory
clean:
  rm -rf build

# Print this help message
help:
  just --list
