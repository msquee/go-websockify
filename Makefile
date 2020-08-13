# Usage:
#
# make
# make clean

.PHONY: all

all: build

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.versionString=$$(git describe --tags --first-parent --abbrev=6 --long --dirty --always) -X main.buildTime=$(date)" -o ./bin/go-websockify

clean:
	rm -rf bin/
