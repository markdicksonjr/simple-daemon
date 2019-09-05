package main

import (
	"flag"
	simple_daemon "github.com/markdicksonjr/simple-daemon"
	"log"
	"os"
)

// in this app, you should run ./main.exe install --name bob
// then, start the service.  You should see "testlogfile" containing the correct response
func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	if err := simple_daemon.Start(simple_daemon.Info{
		Name:         "Test Daemon",
		Description:  "A simple test daemon",
		Dependencies: nil,
	}, simple_daemon.Behavior{
		WorkFn: func() error {
			name := flag.String("name", "", "some name to print")
			flag.Parse()

			if *name != "bob" {
				log.Println("your name should be bob,", *name)
			} else {
				log.Println("simple log for", *name)
			}
			return nil
		},
	}); err != nil {
		log.Fatal(err)
	}
}
