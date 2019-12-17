build:
	@go build -o server server.go

run:
	@./server

clean:
	@rm server
