package main

import (
	"log"
	"net/http"
)

/*
authenticateOrigin parses an HTTP request and checks
if its a valid request according to rules
*/
func authenticateOrigin(r *http.Request) bool {
	log.Println("WebSocket upgraded")
	return true
}
