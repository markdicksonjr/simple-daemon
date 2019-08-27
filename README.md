# Simple Daemon

A super-simple wrapper for making a simple cross-platform service.  With only a few lines of code, you can
run your process as a daemon/service.

## Getting Started

Get this repository:

`go get github.com/markdicksonjr/simple-daemon`

Initialize a simple daemon instance with the basic information about the service (like name and description), along with 
the behavior of the service (most importantly, the function that does the work).

```go
import (
	simple_daemon "github.com/markdicksonjr/simple-daemon"
	"log"
)

func main() {
	if err := simple_daemon.Start(simple_daemon.Info{
		Name:         "Test",
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
```

This will create the service with the name and description provided.  On Windows, the service will be "auto" and will 
run immediately.

## Service Lifecycle

The various lifecycle stages of services can be managed (by default):

`binaryName install`

`binaryName remove`

`binaryName start`

`binaryName stop`

`binaryName status`

To customize these flags, you can handle them yourself and set the "Command" property on the Behavior you pass to the 
service at creation time.

## Restrictions

Due to a third-party implementation detail, service names with spaces will have those spaces converted to underscores.

## Credits

[github.com/takama/daemon](http://github.com/takama/daemon)
