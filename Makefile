# Set the Go executable path (adjust if necessary)
GO=go

# Set the app name and version
APP=eqk
VERSION=1.0.0

# Build the app
build:
	$(GO) build -o $(APP) main.go

# Install the app
install: build
	sudo cp $(APP) /usr/local/bin/

# Clean the build artifacts
clean:
	rm -f $(APP)
