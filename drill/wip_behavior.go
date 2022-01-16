// Configure some struct
package drill

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {
	sys := NewSystem()

	// Configure
	sys.SetHost("localhost")
	sys.SetPort(1899)
	sys.SetDebug(true)

	fmt.Println(sys)
	go sys.Run(context.Background())
	<-time.After(10 * time.Millisecond)
	fmt.Println(sys)
	sys.SetHost("example.com") // noop

}

// NewSystem returns a stopped system
func NewSystem() *System {
	var sys System
	stopped(&sys)
	return &sys
}

type System struct {
	Settings
	Runner

	host  string
	port  int
	debug bool

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
	s.Settings = &disabled{}
	s.Runner = RunFunc(func(context.Context) {})
}

func stopped(s *System) {
	s.state = StateStopped
	s.Settings = &enabled{s}
	s.Runner = RunFunc(s.run)
}

// change Settings behavior by replacing it's implementation

type enabled struct {
	*System
}

func (s *enabled) SetHost(v string) { s.host = v }
func (s *enabled) SetPort(v int)    { s.port = v }
func (s *enabled) SetDebug(v bool)  { s.debug = v }

type disabled struct{}

func (s *disabled) SetHost(v string) { s.warn() }
func (s *disabled) SetPort(v int)    { s.warn() }
func (s *disabled) SetDebug(v bool)  { s.warn() }

func (s *disabled) warn() {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	name := f.Name()
	i := strings.LastIndex(name, ".")
	if i > -1 {
		name = name[i+1:]
	}
	log.Printf("%s: settings disabled", name)
}

type RunFunc func(context.Context)

func (fn RunFunc) Run(ctx context.Context) { fn(ctx) }

type Runner interface {
	Run(context.Context)
}

type Settings interface {
	SetHost(string)
	SetPort(int)
	SetDebug(bool)
}
