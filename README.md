# Go WebSockify

1. Start the test TCP server
Run sudo nc -l -k 1000, this uses netcat to act as the TCP server we are proxying to for testing. must be ran as root so use sudo.

2. Startup Golang proxy and Webserver
You can install Modd (https://github.com/cortesi/modd) to automatically reload on code changes or you can just manually run
go run main.go proxy.go auth.go

3. To start the client
Go into /client and run yarn install and then yarn dev. Or you can use npm install and npm dev if you don't have yarn.

# To Do
* Support for a command line configuration and configuration file (https://github.com/spf13/viper)
* Proxy to UNIX domain sockets
* Wrapping programs (LD_PRELOAD dynamic linking)
* Logging (https://vector.dev)
* SSL (the wss:// WebSockets URI): This is detected automatically by websockify by sniffing the first byte sent from the client and then wrapping the socket if the data starts with '\x16' or '\x80' (indicating SSL).
* OIDC Authentcation and or HTTP Auth and or URI query token auth

# Miscellaneous
* We need a cooler name other than WebSockify, so get thinking!
