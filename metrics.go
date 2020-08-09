package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	bytesTx = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "net",
			Name:      "websocket_bytes_tx_total",
			Help:      "byets sent to tcp",
		})
	bytesRx = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "net",
			Name:      "websocket_bytes_rx_total",
			Help:      "bytes received from tcp",
		})
	wsConnCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "net",
			Name:      "websocket_connections_active",
			Help:      "Active WebSocket connections",
		})
	tcpConnCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "net",
			Name:      "tcp_connections_active",
			Help:      "Active TCP connections",
		})
)

func init() {
	prometheus.MustRegister(bytesTx)
	prometheus.MustRegister(bytesRx)
	prometheus.MustRegister(wsConnCounter)
	prometheus.MustRegister(tcpConnCounter)
}
