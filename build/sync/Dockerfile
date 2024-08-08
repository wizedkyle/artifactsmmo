FROM golang:1.22.5-alpine3.20

WORKDIR /go/src/sync

COPY . /go/src/sync

RUN go get ./...
RUN go build ./cmd/sync

ENTRYPOINT ["./sync"]