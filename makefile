all: clean build

build:
	@go build -o server server.go

run:
	@./server

clean:
	@rm -f server

.PHONY: build clean run all

