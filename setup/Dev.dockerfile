FROM golang:1.24.5-alpine3.22
RUN go telemetry off

RUN apk update && \
	apk add --no-cache curl git make
RUN apk add --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing hurl
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v2.2.2