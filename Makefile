.PHONY: build clean test run

# Set the default target to be executed when make is run without arguments
default: build

# Build the microservice binary and place it in the bin directory
build:
	go build -o bin/auth.exe .

# Clean removes the bin directory and any other generated files
clean:
	rm -rf bin

# Run the development server with hot-reloading using air
dev:
	air -c .air.toml

# Run unit tests
test:
	go test -v ./...

# Run the microservice
run: build
	./bin/auth.exe

# Include additional targets as needed, such as for database migrations, linting, etc.
