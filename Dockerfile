FROM golang:1.19

ADD . /server
WORKDIR /server

RUN go mod tidy
RUN go install 
RUN go install github.com/cosmtrek/air@latest

CMD air --build.cmd "go build -o bin/api ./cmd/server" --build.bin "./bin/api" 
