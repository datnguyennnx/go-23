FROM golang:1.20-alpine
RUN mkdir /build
WORKDIR /build
COPY . .
ENV GOOS=linux CGO_ENABLED=0
RUN set -ex && \
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

RUN go build -o server cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=0 /build/server /usr/bin/

EXPOSE 8080
CMD [ "server" ]