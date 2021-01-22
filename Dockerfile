FROM golang:alpine AS builder
RUN apk add --no-cache git gcc musl-dev

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build .

FROM alpine

LABEL Name=go-echo
LABEL Release=https://github.com/onaci/go-echo
LABEL Url=https://github.com/onaci/go-echo
LABEL Help=https://github.com/onaci/go-echo/issues

LABEL virtual.port=80
LABEL virtual.metrics=2112
LABEL virtual.portal=echo_dev

RUN apk add --no-cache ca-certificates

COPY --from=builder /src/go-echo /opt/go-echo

ENTRYPOINT [ "/opt/go-echo" ]
