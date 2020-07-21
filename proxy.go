package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
)

/*
ProxyServer holds state information about the connection
being proxied
*/
type ProxyServer struct {
	wsConn  *websocket.Conn
	tcpConn *net.TCPConn
	tcpAddr *net.TCPAddr
}

/*
NewWebSocketProxy returns a pointer to a ProxyServer struct
*/
func NewWebSocketProxy(wsConn *websocket.Conn, tcpAddr *net.TCPAddr) *ProxyServer {
	proxyServer := ProxyServer{wsConn, nil, tcpAddr}
	return &proxyServer
}

/*
Start starts the bidirectional communcation channel
between the WebSocket and the remote conection
*/
func (proxyServer *ProxyServer) Start() {
	go proxyServer.webSocketToTCP()
	go proxyServer.tcpToWebSocket()
}

/*
Dial is a function of proxyserver struct that
instantiates a TCP connection to proxyserver.tcpAddr
*/
func (proxyServer *ProxyServer) Dial() error {
	tcpConn, err := net.DialTCP("tcp", nil, proxyServer.tcpAddr)

	if err != nil {
		message := "Dialing fail: " + err.Error()
		log.Println(message)
		_ = proxyServer.wsConn.WriteMessage(websocket.TextMessage, []byte(message))
		return err
	}

	proxyServer.tcpConn = tcpConn

	success := fmt.Sprintf("WebSocket %s connected to %v:%d", proxyServer.wsConn.RemoteAddr(), proxyServer.tcpAddr.IP, proxyServer.tcpAddr.Port)
	log.Println(success)

	return nil
}

func (proxyServer *ProxyServer) tcpToWebSocket() {
	/*
		Infinitely forward TCP data back to WebSocket
	*/
	for {
		buffer := make([]byte, 65536)

		n, err := proxyServer.tcpConn.Read(buffer)
		if err != nil {
			proxyServer.tcpConn.Close()
			break
		}

		log.Println(string([]byte(buffer[0:n])))

		err = proxyServer.wsConn.WriteMessage(websocket.BinaryMessage, buffer[0:n])
		if err != nil {
			log.Println("tcpToWebSocket:", err.Error())
		}
	}
}

func (proxyServer *ProxyServer) webSocketToTCP() {
	for {
		_, data, err := proxyServer.wsConn.ReadMessage()
		if err != nil {
			proxyServer.wsConn.Close()
			proxyServer.tcpConn.Close()
			proxyServer = nil
			break
		}

		_, err = proxyServer.tcpConn.Write(data)
		if err != nil {
			log.Println("webSocketToTCP:", err.Error())
			proxyServer.Dial()
			proxyServer.tcpConn.Write(data)
		}
	}
}
