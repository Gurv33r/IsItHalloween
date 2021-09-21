package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func KeyboardInterruptHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c // receive signal
		log.Println("Entered CRTL+C in Terminal -> Exiting!")
		// put anything else you want to do before exiting here
		os.Exit(0)
	}()
}
