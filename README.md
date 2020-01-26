# www-th3-z-xyz

Source code for [my public homepage](https://www.th3-z.xyz). Features live stream, dynamic anime list, server details, project details, etc.

## Setup

### Requirements

* make
* golang 1.13+
* [fswatch](https://github.com/emcrisostomo/fswatch) - `sudo apt install fswatch`
  - Only required for autoreload (`make serve`)

### Running

Run `make serve` to start the development server on [localhost:5555](http://localhost:5555).

### Building

* Run `make build` to create a release in `./bin`. 
* Run `./bin/server` to start the server on [localhost:5555](http://localhost:5555).


