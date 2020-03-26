FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn,direct

RUN set -ex \
        && apk add --no-cache \
            tzdata ca-certificates upx git \
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .

RUN go mod download

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w"
RUN upx -9 caddy

FROM alpine:latest

LABEL maintainer="shmilwdc <shmilwdc@gmail.com>"

RUN set -ex \
        && apk add --no-cache \
            openssh-client git

COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/release/caddy /usr/local/bin/caddy

RUN caddy -version
RUN caddy -plugins

VOLUME [ "/www","/var/lib/caddy" ]

ENV CADDYPATH /var/lib/caddy

WORKDIR /www

EXPOSE 80 443

ENTRYPOINT [ "caddy" ]