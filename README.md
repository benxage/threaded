# Threaded.chat #

There's three main folders. `Dev/` is build and script related, `client/` is the frontend, and `server/` is all backend related.

**You must follow Go's workspace convention by putting this repo under `$GOPATH/src/github.com/bli940505/` or else dependencies will not resolve. You can manually resolve all the the imports but that will fuck up everyone else.**

## client ##

TODO

## dev ##

Dev contains a script named `dev.sh` that does easy launching for you. Since the front does not yet exist, this just starts the server. 

`dev.sh` can be used in the following ways (you may have to give it permission by running `chmod +x dev.sh`):

* `dev.sh launch` launches the server
* `dev.sh launch background` launches the server in the background without corrupting your current terminal. You may have to manually kill the process later.
* `dev.sh info` calls  `cloc` to parse information about the project. Right now it prints information such as how many lines of code in relation to what language is used.

**`dev.sh` contains Linux support but have not been tested**

## server ##

Server is written with [chi](https://github.com/go-chi/chi), a light weight router that's compatible with Go's native `net/http` library. Database uses [go-pg/pg](https://github.com/go-pg/pg) with ORM support.

To run the server manually, make sure you have [dep](https://github.com/golang/dep) installed, and run `dep ensure` under `server/` to resolve dependencies. Then run `go run server.go`.

### api ###

`server/api/` is where you will write all of your routes. `api.go` is the main entry pointer for the api package.

For every feature you write, define a `my_feature.go` under this directory, and a `RouteMyFeature()` function that returns a `*chi.Mux` object (see [example.go](server/api/example.go)). Then mount this feature as sub-router under `api.go`.

Each route handler you write should take in a `*instance.ServerInstance` parameter if you want to use the development features I have implemented, such as the database or the error handler. 

I would stick with the consistency of using closure to return a `http.HandlerFunc` for your router handlers (see [example.go](server/api/example.go)).

### database ###

`server/database/` is where all of the database code exist. [database.go](server/database/database.go) defines an interface named `Database`, and [postgres.go](server/database/postgres.go) conforms to this interface.

Note `instance.ServerInstance{}` under `server/instance/instance.go` includes a `Database` instance, not a `Postgres` instance. This is for modularity purposes; in the future if we want to switch database we can just add a `NewDatabase.go` and conform it to this interface to be able to use it.

Right now the database is extremely basic. You will probably need to add your own functions. For every function you add, unless it is a database specific operation, please add the function signature under the `Database` interface as well.

**Even if the database operation is simplistic, please still add it under a function wrapper. Again, this is for modularity purposes.

### instance ###

`server/instance/` is where the server instance resides. It includes a config field, a database field, and an error channel. This server instance should be pass around all of the routes for database access, and error handling.

### errors ###

`server/internal/errors/` is where all of the error handling resides. Traditional error handling in Go looks something like this:

```
func foo() error {
    value, err := bar()
    if err != nil {
        // return err or handle the error
    }

    return nil
}
```

There's no `try/catch` block in Go, and it's extremely annoying to do this every time (IMO). As a solution, `instance.ServerInstance{}` contains an error channel. An error handler function named `HandleErrors()` exist under `server/internal/errors/error_handler.go`, which spawns a thread that continuesly listens to the `instance.ServerInstance.Err` channel.

For any error you want to handle, please define it inside `server/internal/errors/error.go` and modify the switch statement inside of `HandleErrors()` to match the behavior you want. There's already some pre-defined errors for you to look at as an example.

To handle errors in this style, pass the error to the error channel (even if it's nil) inside your routes and forget about it. It would look something like this:

```
func MyHandler(in *instance.ServerInstance) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := SomeFunction()
        in.Err <- err

        // do other things
	}
}
```

Unless there's a special case, DO NOT pass the error back up the stack.

### config ###

`server/internal/config/config.go` simple parses a `.toml` file for dynamic configuration purposes. Right now only a `Host` and a `Port` defined.

## Development ##

**DO NOT PUSH TO MASTER!** Don't even try, it's restricted.

Create your own branch with a proper name and submit a PR when you're ready. PRs are configured to required at least one approval before merging.
