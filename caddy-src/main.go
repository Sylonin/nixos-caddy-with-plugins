// This file is copied from:
// https://github.com/caddyserver/caddy/blob/82c356f2548ca62b75f76104bef44915482e8fd9/cmd/caddy/main.go#L21-L25

//  1. Copy this file (main.go) into a new folder
//  2. Edit the imports below to include the modules you want plugged in
//  3. Run `go mod init caddy`
//  4. Run `go install` or `go build` - you now have a custom binary!
//
// Or you can use xcaddy which does it all for you as a command:
// https://github.com/caddyserver/xcaddy
package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/RussellLuo/caddy-ext/ratelimit"
	_ "github.com/RussellLuo/caddy-ext/requestbodyvar"
	_ "github.com/WingLim/caddy-webhook"
	_ "github.com/abiosoft/caddy-exec"
	_ "github.com/abiosoft/caddy-json-parse"
	_ "github.com/aksdb/caddy-cgi/v2"
	_ "github.com/caddy-dns/cloudflare"
	_ "github.com/caddy-dns/namecheap"
	_ "github.com/caddy-dns/porkbun"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/caddyserver/replace-response"
	_ "github.com/caddyserver/transform-encoder"
	_ "github.com/corazawaf/coraza-caddy"
	_ "github.com/cubic3d/caddy-quantity-limiter"
	_ "github.com/git001/caddyv2-upload"
	_ "github.com/greenpau/caddy-appd"
	_ "github.com/greenpau/caddy-git"
	_ "github.com/greenpau/caddy-security"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/http"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/layer4"
	_ "github.com/lolPants/caddy-requestid"
	_ "github.com/mholt/caddy-l4"
	_ "github.com/mpilhlt/caddy-conneg"
	_ "github.com/shift72/caddy-geo-ip"
	_ "github.com/sjtug/caddy2-filter"
	_ "github.com/toowoxx/caddy2-html-injection-plugin"
	_ "github.com/ueffel/caddy-basic-auth-filter"
)

func main() {
	caddycmd.Main()
}
