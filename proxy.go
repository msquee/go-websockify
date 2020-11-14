package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
)

// Proxy interface
type Proxy interface {
	Initialize(*websocket.Conn, *net.TCPAddr) *ProxyServer
	Start()
	Dial() error
	Teardown()
}

// ProxyServer holds state information about the connection
// being proxied.
type ProxyServer struct {
	wsConn  *websocket.Conn
	tcpAddr *net.TCPAddr
	tcpConn *net.TCPConn
}

// Initialize ProxyServer and return struct.
func (p *ProxyServer) Initialize(wsConn *websocket.Conn, tcpAddr *net.TCPAddr) *ProxyServer {
	p.wsConn = wsConn
	p.tcpAddr = tcpAddr

	return p
}

// Start the bidirectional communcation channel
// between the WebSocket and the remote conection.
func (p *ProxyServer) Start() {
	go p.ReadWebSocket()
	go p.ReadTCP()
}

// Dial is a function of proxyserver struct that
// instantiates a TCP connection to proxyserver.tcpAddr
func (p *ProxyServer) Dial() error {
	tcpConn, err := net.DialTCP(p.tcpAddr.Network(), nil, p.tcpAddr)

	if err != nil {
		message := "dialing fail: " + err.Error()
		log.Println(message)

		p.wsConn.WriteMessage(websocket.TextMessage, []byte(message))

		return err
	}

	p.tcpConn = tcpConn
	tcpConnCounter.Inc()

	success := fmt.Sprintf("WebSocket %s connected to %+v:%d", p.wsConn.RemoteAddr(), p.tcpAddr.IP, p.tcpAddr.Port)
	log.Println(success)

	return nil
}

// ReadWebSocket reads from the WebSocket and
// writes to the backend TCP connection.
func (p *ProxyServer) ReadWebSocket() {
	for {
		_, data, err := p.wsConn.ReadMessage()
		if err != nil {
			p.Teardown()
			break
		}

		_, err = p.tcpConn.Write(data)
		if err != nil {
			log.Println("webSocketToTCP:", err.Error())

			p.Dial()
			p.tcpConn.Write(data)
		}

		bytesRx.Add(float64(len(data)))
	}
}

// ReadTCP reads from the backend TCP connection and
// writes to the WebSocket.
func (p *ProxyServer) ReadTCP() {
	buffer := make([]byte, config.bufferSize)

	for {
		bytesRead, err := p.tcpConn.Read(buffer)

		if err != nil {
			p.Teardown()
			break
		}

		if err := p.wsConn.WriteMessage(websocket.BinaryMessage, buffer[:bytesRead]); err != nil {
			log.Println("tcpToWebSocket:", err.Error())
			break
		}

		bytesTx.Add(float64(bytesRead))
	}
}

// Teardown the WebSocket and backend TCP connection.
func (p *ProxyServer) Teardown() {
	p.tcpConn.Close()
	p.wsConn.Close()

	// Decrement Prometheus counters
	tcpConnCounter.Dec()
	wsConnCounter.Dec()
}
