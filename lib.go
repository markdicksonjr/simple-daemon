package simple_daemon

import (
	"github.com/takama/daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Info struct {
	Name         string
	Description  string
	Dependencies []string
}

type Behavior struct {
	WorkFn  func() error
	ExitFn  func() error
	StdLog  *log.Logger
	ErrLog  *log.Logger
	Command string
}

func Start(info Info, behavior Behavior) error {
	srv, err := daemon.New(info.Name, info.Description, info.Dependencies...)
	if err != nil {
		return err
	}

	s := ManagedService{
		Daemon:  srv,
		info:    info,
		workFn:  behavior.WorkFn,
		exitFn:  behavior.ExitFn,
		stdlog:  behavior.StdLog,
		errlog:  behavior.ErrLog,
		command: behavior.Command,
	}

	if s.stdlog == nil {
		s.stdlog = log.New(os.Stdout, "", 0)
	}

	if s.errlog == nil {
		s.errlog = log.New(os.Stderr, "", 0)
	}

	status, err := s.Manage()
	if err != nil {
		return err
	}

	s.stdlog.Println(status)

	return nil
}

type ManagedService struct {
	daemon.Daemon
	workFn  func() error
	exitFn  func() error
	info    Info
	stdlog  *log.Logger
	errlog  *log.Logger
	command string
}

func (s *ManagedService) Manage() (string, error) {
	usage := "Usage: " + s.info.Name + " install | remove | start | stop | status"
	command := s.command

	// if received any kind of command, and we haven't specified one ourselves, do it
	if len(command) == 0 && len(os.Args) > 1 {
		command = os.Args[1]
	}

	if len(command) > 0 {
		switch command {
		case "install":
			return s.Install(os.Args[2:]...)
		case "remove":
			return s.Remove()
		case "start":
			return s.Start()
		case "stop":
			return s.Stop()
		case "status":
			return s.Status()
		case "help":
			return usage, nil
		default:
			break
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	workError := make(chan error, 1)

	go func() {
		if err := s.workFn(); err != nil {
			interrupt <- os.Kill
		}
	}()

	for {
		select {
		case err := <-workError:
			if s.exitFn != nil {
				if exitErr := s.exitFn(); exitErr != nil && s.errlog != nil { // TODO: do something meaningful with this error
					s.errlog.Println(exitErr)
				}
			}
			return "An error occurred", err
		case killSignal := <-interrupt:
			var err error

			if s.exitFn != nil {
				err = s.exitFn()
			}

			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", err
			}
			return "Daemon was killed", err
		}
	}
}
