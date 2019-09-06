package simple_daemon

import "golang.org/x/sys/windows/svc"

func IsInteractive() (bool, error) {
	return svc.IsAnInteractiveSession()
}
