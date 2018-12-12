package errors

import (
	"fmt"
	"os"

	"github.com/bli940505/threaded/server/internal/types"
)

// HandleErrors spawns two thread to listen for system signals and errors.
// Will gracefully stops server on certain signals and errors
func HandleErrors(s *types.Server) {
	shutdown := func() {
		close(s.Err)
		if s.Database != nil {
			fmt.Println("closing database")
			s.Database.Close()
		}
		fmt.Println("shutting down")
		os.Exit(0)
	}

	// spawn error listener
	go func() {
		for err := range s.Err {
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
