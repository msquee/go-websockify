package main

import (
	"log"
	"os"
	"os/signal"
)

// SetupInterruptHandler listens for OS interrupts and shuts down
// Go WebSockify if an interrupt is received.
func SetupInterruptHandler() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		log.Println("Shutting down Go WebSockify")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		} else {
			os.Exit(0)
		}
	}()
}
