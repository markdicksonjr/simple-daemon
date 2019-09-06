package simple_daemon

import "os"

func IsInteractive() (bool, error) {
	return os.Getppid() != 1, nil
}
