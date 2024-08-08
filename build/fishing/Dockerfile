FROM golang:1.22.5-alpine3.20

WORKDIR /go/src/fishing

COPY . /go/src/fishing

RUN go get ./...
RUN go build ./cmd/fishing

ENTRYPOINT ["./fishing"]