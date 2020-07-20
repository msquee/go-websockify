package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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
		Start()
	},
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:    65536,
		WriteBufferSize:   65536,
		CheckOrigin:       authenticateOrigin,
		EnableCompression: true,
	}
	proxyServer *ProxyServer
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", webSocketHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

/*
webSocketHandler handles an incoming HTTP upgrade
request for a WebSocket connection while establishing a
bidirectional stream to a proxied TCP resource.
*/
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New WebSocket Connection from %s", r.RemoteAddr)
	log.Println("Attempting to upgrade WebSocket connection")

	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Failed to upgrade websocket request: ", err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", string("localhost:1000"))
	if err != nil {
		message := "Failed to resolve destination: " + err.Error()
		log.Println(message)
		_ = wsConn.WriteMessage(websocket.CloseMessage, []byte(message))
		return
	}

	proxyServer = NewWebSocketProxy(wsConn, tcpAddr)
	err = proxyServer.Dial()
	if err != nil {
		log.Println(err)
		return
	}
	go proxyServer.Start()
}
