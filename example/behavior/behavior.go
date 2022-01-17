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

func (s *System) switchState(state func(s *System)) {
	s.m.Lock()
	before := s.state
	state(s)
	s.m.Unlock()
	log.Printf("switch state: %s -> %s", before, s.state)
}

type State string

const (
	StateStopped State = "stopped"
	StateRunning State = "running"
)

func running(s *System) {
	s.state = StateRunning
	s.NetSettings = &disabled{}
	s.LogSettings = &disabled{}
	s.Runner = runFunc(func(context.Context) {
		log.Println("already running")
	})
}

func stopped(s *System) {
	s.state = StateStopped
	s.NetSettings = &enabled{s}
	s.LogSettings = &enabled{s}
	s.Runner = runFunc(s.run)
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
