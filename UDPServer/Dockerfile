FROM golang:latest

WORKDIR /UDP_SERVER
COPY . .

RUN go get -u

ENTRYPOINT ["go", "run", "main.go"]