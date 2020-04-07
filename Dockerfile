FROM golang:alpine3.11 AS builder
RUN apk add --no-cache git gcc musl-dev

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build .

FROM alpine:3.11

LABEL Name=go-echo \
      Release=https://github.com/onaci/go-echo \
      Url=https://github.com/onaci/go-echo \
      Help=https://github.com/onaci/go-echo/issues

RUN apk add --no-cache ca-certificates

COPY --from=builder /src/go-echo /opt/go-echo

ENTRYPOINT [ "/opt/go-echo" ]
