## build: build the application and place the built app in the bin folder
build:
	go build -o bin/tblogs ./cmd/tblogs/.

## start: start container
start:
	docker-compose up -d

## test: runs tests
test:
	go test -v ./... --cover

## compile: compiles the application for multiple environments and place the output executables under the bin folder
compile:
	# 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/tblogs-freebsd-64 ./cmd/tblogs/.
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o ./bin/tblogs-macos-64 ./cmd/tblogs/.
	# Linux
	GOOS=linux GOARCH=amd64 go build -o ./bin/tblogs-linux-64 ./cmd/tblogs/.
	# Windows
	GOOS=windows GOARCH=amd64 go build -o ./bin/tblogs-windows-64 ./cmd/tblogs/.

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'