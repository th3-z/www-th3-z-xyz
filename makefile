PID = /tmp/beta-th3-z-xyz.pid
SRC = $(wildcard *.go)
APP = ./server

serve: restart
	@fswatch -m poll_monitor -o . | xargs -n1 -I{} make restart || make kill

kill:
	@-kill `cat $(PID)` || true

before:
	@echo "pre-hook - TODO: Generate static assets"

$(APP): $(SRC)
	@go build -o $@ $?

restart: kill before $(APP)
	@$(APP) & echo $$! > $(PID)

build:
	@go build -o $(APP) server.go
	@rm -rf ./bin
	@mkdir ./bin
	@cp -r ./static ./bin
	@cp -r ./templates ./bin
	@cp ./README.md ./bin
	@cp ./LICENSE ./bin
	@cp ./server ./bin

clean:
	@-kill `cat $(PID)` || true
	@rm -f $(APP)

.PHONY: build clean run all serve restart kill before

