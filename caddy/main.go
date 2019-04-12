package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/caddyserver/dnsproviders/dnspod"
	_ "github.com/linkonoid/caddy-dyndns"
	// _ "github.com/casbin/caddy-authz"
	// _ "github.com/miquella/caddy-awses"
	// _ "github.com/coopernurse/caddy-awslambda"
	// _ "github.com/nicolasazrak/caddy-cache"
	// _ "github.com/jung-kurt/caddy-cgi"
	// _ "github.com/captncraig/cors"
	// _ "github.com/payintech/caddy-datadog"
	// _ "github.com/epicagency/caddy-expires"
	// _ "github.com/echocat/caddy-filter"
	// _ "github.com/caddyserver/forwardproxy"
	_ "github.com/aablinov/caddy-geoip"
	// _ "github.com/abiosoft/caddy-git"
	_ "github.com/awoodbeck/caddy-git"
	// _ "github.com/zikes/gopkg"
	_ "github.com/pieterlouw/caddy-grpc"
	_ "github.com/pyed/ipfilter"
	_ "github.com/BTBurke/caddy-jwt"
	// _ "github.com/simia-tech/caddy-locale"
	_ "github.com/tarent/loginsrv/caddy"
	// _ "github.com/SchumacherFM/mailout"
	// _ "github.com/hacdias/caddy-minify"
	_ "github.com/Xumeiquer/nobots"
	// _ "github.com/miekg/caddy-prometheus"
	_ "github.com/mastercactapus/caddy-proxyprotocol"
	_ "github.com/jung-kurt/caddy-pubsub"
	_ "github.com/xuqingfeng/caddy-rate-limit"
	_ "github.com/captncraig/caddy-realip"
	// _ "github.com/freman/caddy-reauth"
	// _ "github.com/restic/caddy"
	_ "github.com/techknowlogick/caddy-s3browser"
	// _ "github.com/lucaslorentz/caddy-supervisor/httpplugin"
	// _ "go.okkur.org/torproxy"
	_ "github.com/hacdias/caddy-webdav"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}