package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var (
	runAsDaemon bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&runAsDaemon, "daemon", "D", false, "run Go WebSockify as daemon")
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:  "go-websockify",
	Long: `Starts a WebSocket server which facilitates a bidirectional communications channel. Endpoints are responsible for implementing their own transport layer, Go WebSockify's only job is to move buffers from point A to B.`,
	Run: func(cmd *cobra.Command, args []string) {
		if runAsDaemon {
			cntxt := &daemon.Context{
				PidFileName: "go-websockify.pid",
				PidFilePerm: 0644,
			}

			daemon, err := cntxt.Reborn()
			if err != nil {
				log.Fatalf("Unable to start %s", err.Error())
				return
			}
			if daemon != nil {
				return
			}
			defer cntxt.Release()

			log.Println("Starting server listening on port 127.0.0.1:8080")
			daemonMessage := fmt.Sprintf("Running daemon under PID %d", os.Getpid())
			log.Println(daemonMessage)
		} else {
			log.Println("Starting server listening on port 127.0.0.1:8080")
		}
		StartHTTP()
	},
}
