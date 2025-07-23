APP = tblogs

.PHONY: build run test lint clean install help

build:
	go build -o bin/$(APP) ./cmd/$(APP)

run: build
	bin/$(APP)

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/$(APP)

install:
	go install ./cmd/$(APP)

help:
	@echo "Common commands:"
	@echo "  make build    # Build the app"
	@echo "  make run      # Build and run the app"
	@echo "  make test     # Run tests"
	@echo "  make lint     # Run linter"
	@echo "  make clean    # Remove built files"
	@echo "  make install  # Install the app"