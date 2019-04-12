FROM alpine:latest as builder

RUN set -ex \
        && apk add --no-cache \
            tzdata ca-certificates \
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone

FROM frolvlad/alpine-glibc:latest

LABEL maintainer="shmilwdc <shmilwdc@gmail.com>"

ARG TAG="2.0.0b12"
ARG VERSION="v2.0.0-b12"

ENV PORT 8388
ENV PSK=
ENV OBFS http

RUN set -ex \
        && apk add --no-cache \
            libstdc++ \
        && wget -O snell-server.zip https://github.com/surge-networks/snell/releases/download/${TAG}/snell-server-${VERSION}-linux-amd64.zip \
        && unzip snell-server.zip \
        && rm snell-server.zip \
        && mv snell-server /usr/local/bin/

COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY docker-entrypoint.sh /usr/local/bin/

ENTRYPOINT [ "docker-entrypoint.sh" ]