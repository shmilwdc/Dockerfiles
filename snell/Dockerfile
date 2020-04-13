FROM alpine:latest as builder

ENV TAG 2.0.0
ENV VERSION v2.0.0

RUN set -ex \
        && apk add --no-cache \
            tzdata ca-certificates \
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && wget -O snell-server.zip https://github.com/surge-networks/snell/releases/download/${TAG}/snell-server-${VERSION}-linux-amd64.zip \
        && unzip snell-server.zip

FROM frolvlad/alpine-glibc:latest

LABEL maintainer="shmilwdc <shmilwdc@gmail.com>"

ENV PORT 8388
ENV PSK=
ENV OBFS http

RUN set -ex \
        && apk add --no-cache \
            libstdc++

COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /snell-server /usr/local/bin/
COPY docker-entrypoint.sh /usr/local/bin/

ENTRYPOINT [ "docker-entrypoint.sh" ]