FROM alpine:latest as builder

ENV VERSION v4.23.1

RUN set -ex \
        && apk add --no-cache \
            tzdata ca-certificates \
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && wget -O v2ray.zip https://github.com/v2ray/v2ray-core/releases/download/${VERSION}/v2ray-linux-64.zip \
        && unzip v2ray.zip \
        && chmod +x v2ctl v2ray \
        && wget -O geoip.dat https://github.com/v2ray/geoip/releases/latest/download/geoip.dat \
        && wget -O geosite.dat https://github.com/v2ray/domain-list-community/releases/latest/download/dlc.dat

FROM alpine:latest

LABEL maintainer="shmilwdc <shmilwdc@gmail.com>"

ENV PATH /usr/local/bin/v2ray:$PATH

VOLUME /etc/v2ray

COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /v2ray /v2ctl /geoip.dat /geosite.dat /usr/local/bin/v2ray/
COPY --from=builder /config.json /etc/v2ray/

CMD [ "v2ray", "-config", "/etc/v2ray/config.json" ]