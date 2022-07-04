package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  config.bufferSize,
		WriteBufferSize: config.bufferSize,
		CheckOrigin:     authenticateOrigin,
		Subprotocols:    []string{"binary"},
	}
	ctx, stopHTTP = context.WithCancel(context.Background())
	server        = &http.Server{}
)

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

// StartHTTP starts the Go WebSockify web server.
func StartHTTP() {
	defer stopHTTP()

	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{},
	))
	router.HandleFunc(config.httpPath, webSocketHandler)

	server = &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       60 * time.Second,
		Addr:              config.bindAddr,
		Handler:           router,
	}

	if config.cert != "" && config.key != "" {
		log.Print("Certificate and key detected, running HTTPS config for WebSocket server")
		server.TLSConfig = &tls.Config{
			MaxVersion: tls.VersionTLS12,
			Certificates: []tls.Certificate{
				{
					Certificate: [][]byte{
						readFile(config.cert),
					},
					PrivateKey: readFile(config.key),
				},
			},
		}
		server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler))
		listening := fmt.Sprintf("Listening at address %s", config.bindAddr)
		log.Println(listening)
		log.Fatal(server.ListenAndServeTLS(config.cert, config.key))

	} else {
		listening := fmt.Sprintf("Listening at address %s", config.bindAddr)
		log.Println(listening)
		log.Fatal(server.ListenAndServe())
	}

	if ctx.Err() != nil {
		log.Fatalln(ctx.Err())
		return
	}
}

// webSocketHandler handles an incoming HTTP upgrade request
// and starts a bidirectional proxy to the remote connection.
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New WebSocket Connection from %s", r.RemoteAddr)
	log.Println("Attempting to upgrade WebSocket connection")

	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("failed to upgrade websocket request: ", err)
		return
	}
	wsConnCounter.Inc()

	host, port, err := net.SplitHostPort(config.remoteAddr)
	if err != nil {
		log.Println("failed to parse remote address")
		return
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		message := "failed to resolve destination: " + err.Error()
		log.Println(message)
		_ = wsConn.WriteMessage(websocket.CloseMessage, []byte(message))
		return
	}

	var p Proxy = new(ProxyServer)
	p.Initialize(wsConn, tcpAddr)

	if err := p.Dial(); err != nil {
		log.Println(err)
		return
	}

	go p.Start()
}
