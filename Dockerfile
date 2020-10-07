FROM golang:1.14.7 as build

RUN mkdir /go-websockify
WORKDIR /go-websockify
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s -X main.versionString=$(git describe --tags --first-parent --abbrev=6 --long --dirty --always) -X main.buildTime=$(date -Iseconds)" -o /go/bin/go-websockify
FROM scratch
COPY --from=build /go/bin/go-websockify /go/bin/go-websockify
ENTRYPOINT ["/go/bin/go-websockify"]
