FROM alpine:3.12.0
COPY go-websockify /bin/
EXPOSE 1984
CMD [ "/bin/go-websockify", "--bind-addr", "0.0.0.0:8080", "--remote-addr", "127.0.0.1:1984", "--echo" ]
