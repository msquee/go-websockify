package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var (
	runAsDaemon   bool
	printVersion bool
	bindAddress string
	daemonContext daemon.Context
)

func init() {
	SetupInterruptHandler()

	rootCmd.PersistentFlags().BoolVarP(&runAsDaemon, "daemon", "D", false, "run Go WebSockify as daemon")
	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "V", false, "print")
	rootCmd.PersistentFlags().StringVar(&bindAddress, "bind-addr", "", "Bind server to port")


}

func main() {
	Execute()
}
//Execute runs the default command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

var rootCmd = &cobra.Command{
	Use:  "go-websockify",
	Long: `Starts a WebSocket server which facilitates a bidirectional communications channel. Endpoints are responsible for implementing their own transport layer, Go WebSockify's only job is to move buffers from point A to B.`,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println("1.0.0")
			os.Exit(0)
		}
		
		if bindAddress != ""{
			host, port, err := net.SplitHostPort(bindAddress)

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			//Check if valid host
			ip := net.ParseIP(host)
			if ip == nil{
				log.Println("Invalid Host!")
				os.Exit(1)
			}

			//Check if valid prot
			validPort := regexp.MustCompile("^0*(?:6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$")
			if !validPort.MatchString(port){
				log.Println("Invalid Port!")
				os.Exit(1)
			}

		}

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
		
		log.Println("Starting Go WebSockify")
		StartHTTP(bindAddress)
	},
}
