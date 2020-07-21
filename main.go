package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var (
	runAsDaemon   bool
	daemonContext daemon.Context
)

func init() {
	SetupInterruptHandler()

	rootCmd.PersistentFlags().BoolVarP(&runAsDaemon, "daemon", "D", false, "run Go WebSockify as daemon")
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

var rootCmd = &cobra.Command{
	Use:  "go-websockify",
	Long: `Starts a WebSocket server which facilitates a bidirectional communications channel. Endpoints are responsible for implementing their own transport layer, Go WebSockify's only job is to move buffers from point A to B.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Go WebSockify")

		if runAsDaemon {
			daemonContext := &daemon.Context{
				PidFileName: "go-websockify.pid",
				PidFilePerm: 0644,
			}

			daemon, err := daemonContext.Reborn()
			if err != nil {
				log.Fatalf("Unable to start %s", err.Error())
				return
			}
			if daemon != nil {
				return
			}
			defer daemonContext.Release()

			daemonMessage := fmt.Sprintf("Running Go WebSockify daemon under PID %d", os.Getpid())
			log.Println(daemonMessage)
		}
		StartHTTP()
	},
}
