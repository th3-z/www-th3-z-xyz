build:
	@cd ./ui; \
	ng build
	@cd server; \
	go build server.go

test:
	@cd ./ui; \
	ng test

e2e:
	@cd ./ui; \
	ng e2e

run:
	@cd server;\
	./server

