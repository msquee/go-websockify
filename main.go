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
	bindAddr    string
	remoteAddr  string

	daemonContext daemon.Context
)

func init() {
	SetupInterruptHandler()

	rootCmd.PersistentFlags().StringVar(&bindAddr, "bind-addr", "localhost:8080", "Bind address")
	rootCmd.PersistentFlags().StringVar(&remoteAddr, "remote-addr", ":3000", "Remote address")
	rootCmd.PersistentFlags().BoolVarP(&runAsDaemon, "daemon", "D", false, "Run Go WebSockify as daemon")

	rootCmd.MarkPersistentFlagRequired("remote-addr")
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
	TraverseChildren: true,
	Use:              "go-websockify",
	Long:             `Starts a WebSocket server which facilitates a bidirectional communications channel. Endpoints are responsible for implementing their own transport layer, Go WebSockify's only job is to move buffers from point A to B.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Go WebSockify")

		if runAsDaemon {
			log.Println("Running Go WebSockify as daemon")

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

			daemonMessage := fmt.Sprintf("Daemon running under PID %d", os.Getpid())
			log.Println(daemonMessage)
		}
		StartHTTP()
	},
}
