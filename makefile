PID = /tmp/www-th3-z-xyz.pid
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

clean:
	@kill `cat $(PID)` || true
	@rm -f $(APP)
	@rm -f storage.db
	@rm -f static/pastes/*

.PHONY: build clean run all serve restart kill before

