package behavior

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// NewSystem returns a stopped system
func NewSystem() *System {
	var sys System
	stopped(&sys)
	return &sys
}

type System struct {
	NetSettings // controls e.g. host and port
	host        string
	port        int

	LogSettings // e.g. debug mode
	debug       bool

	Runner
	m     sync.Mutex // protects system state switching
	state State
}

func (me *System) String() string {
	return fmt.Sprintf("%s %s:%v", me.state, me.host, me.port)
}

func (me *System) Is(v State) bool {
	return me.state == v
}

func (s *System) run(ctx context.Context) {
	log.Print("run")
	s.switchState(running)

	for {
		select {
		case <-ctx.Done():
			s.switchState(stopped)
			return
		case <-time.After(time.Second):
		}
	}
}

func (s *System) switchState(state func(s *System) (State, error)) {
	s.m.Lock()
	before := s.state
	newState, err := state(s)
	if err == nil {
		s.state = newState
	}
	s.m.Unlock()

	var msg string
	if err != nil {
		msg = " " + err.Error()
	}
	log.Printf("switchState: %s -> %s%s", before, s.state, msg)
}

type State string

const (
	StateStopped State = "stopped"
	StateRunning State = "running"
)

func running(s *System) (State, error) {
	s.NetSettings = &disabled{}
	s.LogSettings = &disabled{}
	s.Runner = runFunc(func(context.Context) {
		log.Println("already running")
	})
	return StateRunning, nil
}

func stopped(s *System) (State, error) {
	s.NetSettings = &enabled{s}
	s.LogSettings = &enabled{s}
	s.Runner = runFunc(s.run)
	return StateStopped, nil
}

// change Settings behavior by replacing it's implementation

type enabled struct {
	*System
}

func (s *enabled) SetHost(v string) { s.host = v }
func (s *enabled) SetPort(v int)    { s.port = v }
func (s *enabled) SetDebug(v bool)  { s.debug = v }

type disabled struct{}

func (s *disabled) SetHost(v string) { log.Print("SetHost: disabled") }
func (s *disabled) SetPort(v int)    { log.Print("SetPort: disabled") }
func (s *disabled) SetDebug(v bool)  { log.Print("SetDebug: disabled") }

type runFunc func(context.Context)

func (fn runFunc) Run(ctx context.Context) { fn(ctx) }

type Runner interface {
	Run(context.Context)
}

type NetSettings interface {
	SetHost(string)
	SetPort(int)
}

type LogSettings interface {
	SetDebug(bool)
}
