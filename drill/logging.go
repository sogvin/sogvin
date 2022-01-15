// Basic use of log printer fucs
package drill

import (
	"log"
)

func init() {
	log.Println("Application", "start")
	log.Print("no", "space")
	log.Printf("Hello, %s!", "world")
	log.Fatal("stop application")
}
