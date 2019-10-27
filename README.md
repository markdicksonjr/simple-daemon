# Simple Daemon

A super-simple wrapper for making a simple cross-platform service.  With only a few lines of code, you can install and run your process as a daemon/service.

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
        UseExeDirAsCwd: true,
	}); err != nil {
		log.Fatal(err)
	}
}
```

This will create the service with the name and description provided.  On Windows, the service will be "auto" and will 
run immediately.  The "UseExeDirAsCwd" flag determines whether or not the working directory is set to wherever the exe
is run from.  Typically, if this is false on windows, this means the current directory will be something like
C:/Windows/system32 or such.

## Service Lifecycle

The various lifecycle stages of services can be managed (by default):

`binaryName install` or `binaryName install [args...]`

`binaryName uninstall`

`binaryName start`

`binaryName stop`

`binaryName status`

For direct invocation:

`binaryName ...otherArgs`

## Credits

[github.com/kardianos/service](https://github.com/kardianos/service)
