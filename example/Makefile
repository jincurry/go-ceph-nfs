# Set Go compiler
GO := go

# Set binary file name
BINARY_NAME := go_ceph_nfs

# Set source file directory
SRC_DIR := .

# Set build flags and compile options
LDFLAGS := -s -w

.PHONY: all build clean install

# Default target
all: build

# Build target
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) $(SRC_DIR)

# Clean target
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

# Install target
install:
	$(GO) install

# Run tests
test:
	$(GO) test ./...
