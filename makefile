all: clean build

build:
	@go build -o server server.go

run:
	@./server

clean:
	@rm server

.PHONY: build clean run all

