# Print this help message
default:
  just --list

# Build the draftsman binary
build:
  mkdir -p build
  CGO_ENABLED=false go build -o build/draftsman main.go

# Run draftsman 
run:
  go run main.go

# Run the tests
test:
  go test -v ./...

# Run the tests with coverage
coverage:
  mkdir -p build
  go test -coverprofile=coverage.out ./...
  go tool cover -html=build/coverage.out

# Run the linter
lint:
  go vet ./...

# Clean the build directory
clean:
  rm -rf build

# Print this help message
help:
  just --list
