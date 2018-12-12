package errors

import (
	"fmt"
	"os"

	"github.com/bli940505/threaded/server/instance"
)

// HandleErrors spawns two thread to listen for system signals and errors.
// Will gracefully stops server on certain signals and errors
func HandleErrors(in *instance.ServerInstance) {
	shutdown := func() {
		close(in.Err)
		if in.Database != nil {
			fmt.Println("closing database")
			in.Database.Close()
		}
		fmt.Println("shutting down")
		os.Exit(0)
	}

	// spawn error listener
	go func() {
		for err := range in.Err {
			switch err {
			case nil:
				continue
			default:
				fmt.Printf("caught error: %+v\n", err)
				shutdown()
			}
		}
	}()
}
