package simple_daemon

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

type Info struct {
	Name         string
	Description  string
	Dependencies []string
}

type Behavior struct {
	WorkFn func() error
	ExitFn func() error
}

type program struct {
	B Behavior
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	if err := p.B.WorkFn(); err != nil {
		log.Fatal(err) // todo
	}
}

func (p *program) Stop(s service.Service) error {
	if p.B.ExitFn != nil {
		if err := p.B.ExitFn(); err != nil {
			return err
		}
	}
	return nil
}

func Start(info Info, behavior Behavior) error {
	svcConfig := &service.Config{
		Name:        info.Name,
		DisplayName: info.Name,
		Description: info.Description,
		Arguments:   nil,
	}

	if len(os.Args) > 2 && os.Args[1] == "install" {
		svcConfig.Arguments = os.Args[2:]
	}

	prg := &program{
		B: behavior,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			return s.Install()
		} else if os.Args[1] == "uninstall" {
			return s.Uninstall()
		} else if os.Args[1] == "start" {
			return s.Start()
		} else if os.Args[1] == "stop" {
			return s.Stop()
		} else if os.Args[1] == "status" {
			status, err := s.Status()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(status)
			return nil
		}
	}

	// TODO: improve loggers

	if err = s.Run(); err != nil {
		return err
	}
	return nil
}
