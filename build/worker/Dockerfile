FROM golang:1.22.5-alpine3.20

WORKDIR /go/src/worker

COPY . /go/src/worker

RUN go get ./...
RUN go build ./cmd/worker

ENTRYPOINT ["./worker"]