package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"

	"github.com/msquee/go-websockify/util"
)

var config struct {
	bindAddr   string
	remoteAddr string
	bufferSize int
	httpPath   string

	runAsDaemon bool
	showVersion bool
	echoServer  bool

	versionString string
	buildTime     string
}
var daemonContext daemon.Context

func init() {
	SetupInterruptHandler()

	rootCmd.PersistentFlags().StringVar(&config.bindAddr, "bind-addr", "0.0.0.0:8080", "bind address")
	rootCmd.PersistentFlags().StringVar(&config.remoteAddr, "remote-addr", "127.0.0.1:1984", "remote address")
	rootCmd.PersistentFlags().IntVar(&config.bufferSize, "buffer", 65536, "buffer size")
	rootCmd.PersistentFlags().BoolVar(&config.echoServer, "echo", false, "sidecar echo server")
	rootCmd.PersistentFlags().StringVar(&config.httpPath, "path", "/websockify", "url path clients connect to")

	rootCmd.Flags().BoolVarP(&config.runAsDaemon, "daemon", "D", false, "run as daemon")
	rootCmd.Flags().BoolVarP(&config.showVersion, "version", "v", false, "print version")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	TraverseChildren: true,
	Use:              "go-websockify",
	Long:             `Starts a TCP to WebSocket proxy.`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.showVersion {
			fmt.Println(fmt.Sprintf("Go WebSockify version %s built on %s", config.versionString, config.buildTime))
			os.Exit(0)
		}

		log.Println("Starting Go WebSockify")

		if config.echoServer {
			go util.StartEchoTCPServer()
		}

		if config.runAsDaemon {
			log.Println("Running Go WebSockify as daemon")

			daemonContext := &daemon.Context{
				PidFileName: "go-websockify.pid",
				LogFileName: "go-websockify.log",
				PidFilePerm: 0644,
				LogFilePerm: 0644,
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
