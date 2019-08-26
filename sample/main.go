package main

import (
	simple_daemon "github.com/markdicksonjr/simple-daemon"
	"log"
)

func main() {
	if err := simple_daemon.Start(simple_daemon.Info{
		Name:         "Test Daemon",
		Description:  "A simple test daemon",
		Dependencies: nil,
	}, simple_daemon.Behavior{
		WorkFn: func() error {
			log.Println("simple log")
			return nil
		},
	}); err != nil {
		log.Fatal(err)
	}
}