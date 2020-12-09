# Usage:
#
# make
# make clean


PLATFORM=linux-amd64

parseOS=$(firstword $(subst -, ,$1))
parseArch=$(or $(word 2,$(subst -, ,$1)),$(value 2))

OS=$(call parseOS, ${PLATFORM})
ARCH=$(call parseArch, ${PLATFORM})

.PHONY: all

all: build

build:
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags="-w -s -X main.versionString=$$(git describe --tags --first-parent --abbrev=6 --long --dirty --always) -X main.buildTime=$(date)" -o ./bin/go-websockify-${OS}-${ARCH}

clean:
	rm -rf bin/
