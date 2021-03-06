package errors

import (
	"log"
	"os"

	"github.com/bli940505/threaded/server/instance"
)

// HandleErrors spawns a thread to listen for the error channel.
// Will gracefully stops server on certain errors
func HandleErrors(in *instance.ServerInstance) {
	shutdown := func() {
		close(in.Err)
		if in.Database != nil {
			log.Println("closing database")
			in.Database.Close()
		}
		log.Println("shutting down")
		os.Exit(0)
	}

	// spawn error listener
	go func() {
		for err := range in.Err {
			switch err {
			case nil:
				continue
			case FatalError, InternalServerError:
				log.Printf("caught error: %+v\n", err)
				shutdown()
			default:
				log.Printf("caught error: %+v\n", err)
				continue
			}
		}
	}()
}
