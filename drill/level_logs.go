// Simple level loggers
//
// When you want to control amount of log information, loggers of
// different levels can be used.
package drill

import (
	"log"
	"os"
)

func init() {
	// default logger is used for INFO level
	log.SetPrefix("INFO ")

	// Move level prefix after timestamp
	flags := log.LstdFlags | log.Lmsgprefix
	log.SetFlags(flags)

	// debug level logger, Note! both loggers use os.Stderr, which
	// could result in competing writes
	debugLog := log.New(os.Stderr, "DEBUG ", flags)

	log.Println("Application start")
	debugLog.Println("somethig happended")
}
