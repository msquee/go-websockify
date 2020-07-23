const { Terminal } = require("xterm")
import "~node_modules/xterm/css/xterm.css"

const ws = new WebSocket('ws://localhost:8080/ws', ['binary'])

const term = new Terminal()
term.open(document.getElementById('terminal'))

console.log(ws.binaryType)

ws.addEventListener('open', () => {
  term.write('Connected to \x1B[1;3;31mGo WebSockify\x1B[0m')
  ws.send(`${Date.now()} web socket message: connected!\u001b[32;1m\n`)
})

ws.addEventListener('message', (frame) => {
  let data = frame.data
  let dataView = new DataView(data)

  console.log(dataView)

  term.write(data.data)
})

ws.addEventListener('error', (e) => {
  console.error(e)
})

document.getElementById('sendMessage').addEventListener('click', () => {
  ws.send(`${Date.now()} clicked button send message!\n`)
})
