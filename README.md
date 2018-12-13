# Threaded.chat #

There's three main folders. `Dev/` is build and script related, `client/` is the frontend, and `server/` is all backend related.

## client ##

TODO

## dev ##

Dev contains a script named `dev.sh` that does easy launching for you. Since the front does not yet exist, this just starts the server. 

`dev.sh` can be used in the following ways:

* `dev.sh launch` launches the server
* `dev.sh launch background` launches the server in the background without corrupting your current terminal
* `dev.sh info` calls  `cloc` to parse information about the project. Right now it prints information such as how many lines of code in relation to what language is used.

**`dev.sh` contains Linux support but have not been tested**

## server ##

Server is written with [chi](https://github.com/go-chi/chi), a light weight router that's compatible with Go's native `net/http` library. Database uses [go-pg/pg](https://github.com/go-pg/pg) with ORM support.

### api ###

`server/api/` is where you will write all of your routes. `api.go` is the main entry pointer for the api package.

For every feature you write, define a `my_feature.go` under this directory, and a `RouteMyFeature()` function that returns a `*chi.Mux` object (see [example.go](server/api/example.go)). Then mount this feature as sub-router under `api.go`.

Each route handler you write should take in a `*instance.ServerInstance` parameter if you want to use the development features I have implemented, such as the database or the error handler. 

I would stick with the consistency of using closure to return a `http.HandlerFunc` for your router handlers (see [example.go](server/api/example.go)).

### database ###

`server/database/` is where all of the database code exist. [database.go](server/database/database.go) defines an interface named `Database`, and [postgres.go](server/database/postgres.go) conforms to this interface. Note `instance.ServerInstance` under `server/instance/instance.go` includes a 