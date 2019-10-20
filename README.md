# beta.th3-z.xyz

A rewrite of www-th3-z-xyz in Golang+Angular.

## Setup

### Requirements

* `go get github.com/labstack/echo` - Echo
* `go get github.com/mattn/go-sqlite3` - SQLite3 driver
* `sudo apt install npm` - NPM
* `cd ui; npm install` - Angular dependencies

### Running

Run `go run server/server.go` to start the backend on 5555

Run `ng serve` from the `ui` folder to start the frontend on 4200

## Tests

Run `ng test` from the `ui` to execute the unit tests via
[Karma](https://karma-runner.github.io).

Run `ng e2e` from the `ui` to execute the end-to-end tests via
[Protractor](http://www.protractortest.org/).

## Building

Run `go build server/server.go` to build the backend.

Run `ng build` from the `ui` folder to build the frontend. The build artifacts
will be stored in the `dist/` directory. Use the `--prod` flag for a 
production build.
