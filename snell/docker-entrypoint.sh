#!/bin/sh

if [ -z "$PSK" ]; then
  PSK=`hexdump -n 16 -e '4/4 "%08x" 1 "\n"' /dev/urandom`
fi

cat > snell-server.conf << EOF
[snell-server]
listen = 0.0.0.0:${PORT}
psk = ${PSK}
obfs = ${OBFS}
EOF

cat snell-server.conf

if [ -z "$@" ]; then
  snell-server
fi

exec "$@"