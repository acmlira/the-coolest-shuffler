FROM golang:1.18-alpine as builder

WORKDIR /build

ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download -x
COPY ./ ./
RUN apk add --no-cache make

ADD build/package/entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]
