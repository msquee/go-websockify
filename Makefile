# Usage:
#
# make
# make clean
# make docker

.PHONY: all

all: build

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.versionString=$$(git describe --tags --first-parent --abbrev=6 --long --dirty --always) -X main.buildTime=$(date)" -o ./bin/go-websockify

docker:
	CGO_ENABLED=0 make build
	cp ./bin/go-websockify ./dev/app
	make clean
	cd dev && docker-compose up --build

clean:
	rm -rf bin/
