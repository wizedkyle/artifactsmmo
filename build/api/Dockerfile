FROM golang:1.22.5-alpine3.20

ENV PORT=9000

WORKDIR /go/src/api

COPY . /go/src/api

RUN go get ./...
RUN go build ./cmd/api

EXPOSE $PORT

ENTRYPOINT ["./api"]