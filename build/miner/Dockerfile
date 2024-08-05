FROM golang:1.22.5-alpine3.20

WORKDIR /go/src/miner

COPY . /go/src/miner

RUN go get ./...
RUN go build ./cmd/miner

ENTRYPOINT ["./miner"]