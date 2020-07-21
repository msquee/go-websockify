package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  65536,
		WriteBufferSize: 65536,
		CheckOrigin:     authenticateOrigin,
	}
	proxyServer   *ProxyServer
	ctx, stopHTTP = context.WithTimeout(context.Background(), time.Second)
	server        = &http.Server{}
)

func StartHTTP() {
	defer stopHTTP()

	router := mux.NewRouter()
	router.HandleFunc("/ws", webSocketHandler)

	server = &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       60 * time.Second,
		Addr:              "localhost:8080",
		Handler:           router,
	}

	log.Println("Listening at address 127.0.0.1:8080")
	log.Fatal(server.ListenAndServe())

	if ctx.Err() != nil {
		log.Fatalln(ctx.Err())
	}
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
