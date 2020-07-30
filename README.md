# Go WebSockify
> RFC 6455 compliant TCP and Unix socket proxy to WebSockets.

Go WebSockify is a pure Go implementation of [novnc/websockify](https://github.com/novnc/websockify) TCP/Unix to WebSocket proxy with improved connection handling. Runs on Linux, Windows and MacOS.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
  - [Quick installation]()
  - [Using prebuilt binaries]()
  - [Building from source]()
- [Usage](#usage)
- [Contributing](#contributing)
  - [Roadmap](#roadmap)
- [Development](#development)
- [License](#license)

## Overview
`TODO`

## Installation
`TODO`

### Quick installation
`TODO`

### Using prebuilt images
`TODO`

### Building from source
`TODO`

## Roadmap
`TODO`

## Usage
```shell
$ go-websockify --help
Starts a WebSocket server which facilitates a bidirectional communications channel. Endpoints are responsible for implementing their own transport layer, Go WebSockify's only job is to move buffers from point A to B.

Usage:
  go-websockify [flags]

Flags:
      --bind-addr string     bind address (default "0.0.0.0:8080")
      --buffer int           buffer size (default 65536)
  -D, --daemon               run Go WebSockify as daemon
  -h, --help                 help for go-websockify
      --remote-addr string   remote address (default "127.0.0.1:3000")
  -v, --version              print Go WebSockify version
```

## Contributing
Both pull requests and issues are welcome on [GitHub](https://github.com/msquee/go-websockify). Look at [CONTRIBUTING.md](https://github.com/msquee/go-websockify/blob/master/CONTRIBUTING.md) to learn more about the coding standards that are enforced on pull requests.

## Development
Instructions for development are located in [CONTRIBUTING.md](https://github.com/msquee/go-websockify/blob/master/CONTRIBUTING.md).

## License
This project is licensed under the terms of the [MIT License](https://github.com/msquee/go-websockify/blob/master/LICENSE.md).