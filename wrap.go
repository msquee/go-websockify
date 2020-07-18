package main

import (
	"log"
	"os"
	"os/exec"
)

/*
WrapProcess does what's on the tin
*/
func WrapProcess(cmd string, args []string) {
	log.Println("Wrapping process")
	wrapped := exec.Command(cmd, args...)
	wrapped.Env = append(os.Environ(),
		"LD_PRELOAD=/usr/lib/test.so",
	)

	cmdOut, err := wrapped.Output()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(cmdOut))
}
