package main

import (
	"log"
	"net"
	"net/http"

	"./cmd"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:    65536,
		WriteBufferSize:   65536,
		CheckOrigin:       authenticateOrigin,
		EnableCompression: true,
	}
	proxyServer *ProxyServer
)

func init() {
	log.Println("Starting Go WebSockify")
}

func main() {
	cmd.Execute()
	log.Println("ASDAJSDLAJSDSAD")
	//startWebSockify()
}

/*
	StartWebSockify set up config
*/
func startWebSockify() {
	log.Println("Starting")
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
