FROM golang:alpine as builder
RUN set -ex \
        && apk add --no-cache git \
        && go get github.com/shadowsocks/v2ray-plugin

FROM alpine:latest

LABEL maintainer="shmilwdc <shmilwdc@gmail.com>"

ENV SERVER_ADDR         0.0.0.0
ENV SERVER_PORT         8388
ENV PASSWORD            password
ENV METHOD              chacha20-ietf-poly1305
ENV TIMEOUT             300
ENV DNS_ADDRS           8.8.8.8,8.8.4.4
ENV ARGS=

USER root

RUN set -ex \
        && apk add --no-cache --virtual .build-deps \
                git \
                autoconf \
                automake \
                build-base \
                c-ares-dev \
                libev-dev \
                libtool \
                libsodium-dev \
                linux-headers \
                mbedtls-dev \
                pcre-dev \
        && cd /tmp \
        && (git clone https://github.com/shadowsocks/shadowsocks-libev.git \
        && cd shadowsocks-libev \
        && git submodule update --init --recursive \
        && ./autogen.sh \
        && ./configure --prefix=/usr --disable-documentation \
        && make install) \
        && (git clone https://github.com/shadowsocks/simple-obfs.git \
        && cd simple-obfs \
        && git submodule update --init --recursive \
        && ./autogen.sh \
        && ./configure --prefix=/usr --disable-documentation \
        && make install) \
        && apk del .build-deps \
        && apk add --no-cache \
                rng-tools \
                ca-certificates \
                $(scanelf --needed --nobanner /usr/bin/ss-* \
                | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
                | sort -u) \
        && rm -rf /tmp/*

COPY --from=builder /go/bin/v2ray-plugin /usr/bin

USER nobody

CMD exec ss-server \
        -s $SERVER_ADDR \
        -p $SERVER_PORT \
        -k $PASSWORD \
        -m $METHOD \
        -t $TIMEOUT \
        -d $DNS_ADDRS \
        --fast-open \
        $ARGS