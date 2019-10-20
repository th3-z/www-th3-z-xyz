# beta.th3-z.xyz

A rewrite of www-th3-z-xyz in Golang+Angular.

## Setup

### Requirements

* `sudo apt install make` - Make
* `go get github.com/labstack/echo` - Echo
* `go get github.com/mattn/go-sqlite3` - SQLite3 driver
* `sudo apt install npm` - NPM
* `cd ui; npm install` - Angular dependencies

## Building

Run `make build`.

### Running

Run `make run` to start the server on 5555

## Tests

Run `make test` to execute the unit tests via
[Karma](https://karma-runner.github.io).

Run `make e2e` to execute the end-to-end tests via
[Protractor](http://www.protractortest.org/).


