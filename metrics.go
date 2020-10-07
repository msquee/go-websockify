package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	bytesTx = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "go_websockify",
			Name:      "websocket_bytes_tx_total",
		},
	)

	bytesRx = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "go_websockify",
			Name:      "websocket_bytes_rx_total",
		},
	)

	wsConnCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "go_websockify",
			Name:      "websocket_connections_active",
			Help:      "Active WebSocket connections",
		})

	tcpConnCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "go_websockify",
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
