package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sogvin/sogvin/example/behavior"
)

func main() {
	sys := behavior.NewSystem()

	// Configure is the act of using the system settings
	sys.SetHost("localhost")
	sys.SetPort(1899)
	sys.SetDebug(true)

	fmt.Println(sys)
	go sys.Run(context.Background())
	<-time.After(10 * time.Millisecond)
	fmt.Println(sys)
	sys.SetHost("example.com") // noop

	sys.Run(context.Background())
}
