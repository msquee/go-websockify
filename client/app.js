import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import fontfaceobserver from 'fontfaceobserver'
const prettyBytes = require('pretty-bytes')

new fontfaceobserver('Iosevka').load().then(() => {
  terminal.open(document.getElementById('terminal'))
  fitAddon.fit()
})

const terminal = new Terminal({
  fontSize: 16,
  fontFamily: 'Iosevka',
  letterSpacing: 0,
}) 
const fitAddon = new FitAddon()

const WS_URL = 'ws://localhost:8080/ws'
let closing = false

let ws = new WebSocket(WS_URL, ['binary'])
ws.binaryType = 'blob'
let counters = {
  tx: 0,
  rx: 0,
}

document.addEventListener('DOMContentLoaded', function () {
  writeMessage('Initiating connection...')
  setupWsHandlers(ws)

  document.getElementById('ping').addEventListener('click', () => {
    if (ws.readyState == ws.OPEN) {
      sendWsMessage('PONG')
    }
  })

  document.getElementById('benchmark').addEventListener('click', () => {
    if (ws.readyState == ws.OPEN) {
      setInterval(function() {
        let buf = Buffer.alloc(Math.floor(Math.random() * 35325) + 1)
        sendWsMessage(buf)
        buf = null
      }, 1000)
    }
  })

  document.getElementById('reconnect').addEventListener('click', () => {
    if (!closing) {
      closing = true
      ws.close(1000)
      if (ws.readyState == ws.CLOSING || ws.CLOSED)
        writeMessage('Disconnected from Go WebSockify')
      ws = new WebSocket(WS_URL, ['binary'])
      setupWsHandlers(ws)
    }
  })
})

function writeMessage(message) {
  terminal.writeln(`${new Date().toISOString()}: ${message}`)
}

function updateCounters(tx = 0, rx = 0) {
  counters.tx += tx
  counters.rx += rx

  document.getElementById('tx').innerText =
    'Transmit ' + prettyBytes(counters.tx)
  document.getElementById('rx').innerText =
    'Receive ' + prettyBytes(counters.rx)
}

function sendWsMessage(message) {
  ws.send(message)
  updateCounters(Buffer.from(message).length, 0)
}

function setupWsHandlers(ws) {
  ws.addEventListener('open', () => {
    closing = false
    writeMessage('Connected to Go WebSockify')
    sendWsMessage('Hello from the browser')
  })

  ws.onmessage = function (frame) {
    const { type, data } = frame
    updateCounters(0, data.size)

    if (data instanceof Blob) {
      data.arrayBuffer().then((buf) => {
        let view = new DataView(buf, 0)
        let text = new TextDecoder('utf-8')
        writeMessage(`Echo: Received frame ${prettyBytes(view.byteLength)}`)
      })
    } else {
      writeMessage(data)
    }
  }

  ws.onerror = function (e) {
    if (ws.readyState == ws.CLOSED) {
      writeMessage('Failed to communicate, is Go WebSockify running?')
    } else {
      writeMessage(`Unknown error: ${JSON.stringify(e)}`)
    }
  }
}
