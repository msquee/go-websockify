package main

import (
	"log"
	"os"
	"os/signal"
)

var (
	signals = make(chan os.Signal, 1)
)

func SetupInterruptHandler() {
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
