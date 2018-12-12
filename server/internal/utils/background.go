package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bli940505/threaded/server/internal/types"
)

// BackgroundHandler spawns two thread to listen for system signals and errors.
// Will gracefully stops server on certain signals and errors
func BackgroundHandler(s *types.Server) {
	// rely all system signals
	signal.Notify(s.Sigs)

	// spawn signal listener
	go func() {
		for sigs := range s.Sigs {
			fmt.Printf("caught signal: %+v\n", sigs)
			switch sigs {
			case syscall.SIGTERM:
				fmt.Println("shutting down")
				close(s.Err)
				close(s.Sigs)
				if s.Database != nil {
					s.Database.Close()
				}
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("shutting down")
				close(s.Err)
				close(s.Sigs)
				if s.Database != nil {
					s.Database.Close()
				}
				os.Exit(0)
			default:
				fmt.Println("default signal")
				continue
			}
		}
	}()

	// spawn error listener
	go func() {
		for err := range s.Err {
			switch err {
			case nil:
				continue
			default:
				fmt.Printf("caught error: %+v\n", err)
				fmt.Println("shutting down")
				close(s.Sigs)
				close(s.Err)
				if s.Database != nil {
					s.Database.Close()
				}
				os.Exit(0)
			}
		}
	}()
}
