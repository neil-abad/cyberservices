package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "cyberservices.com/core/app/dispatcher"
	_ "cyberservices.com/core/app/proxyman/inbound"
	_ "cyberservices.com/core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "cyberservices.com/core/app/commander"
	_ "cyberservices.com/core/app/log/command"
	_ "cyberservices.com/core/app/proxyman/command"
	_ "cyberservices.com/core/app/stats/command"

	// Other optional features.
	_ "cyberservices.com/core/app/dns"
	_ "cyberservices.com/core/app/log"
	_ "cyberservices.com/core/app/policy"
	_ "cyberservices.com/core/app/reverse"
	_ "cyberservices.com/core/app/router"
	_ "cyberservices.com/core/app/stats"

	// Inbound and outbound proxies.
	_ "cyberservices.com/core/proxy/blackhole"
	_ "cyberservices.com/core/proxy/dns"
	_ "cyberservices.com/core/proxy/dokodemo"
	_ "cyberservices.com/core/proxy/freedom"
	_ "cyberservices.com/core/proxy/http"
	_ "cyberservices.com/core/proxy/mtproto"
	_ "cyberservices.com/core/proxy/shadowsocks"
	_ "cyberservices.com/core/proxy/socks"
	_ "cyberservices.com/core/proxy/trojan"
	_ "cyberservices.com/core/proxy/vless/inbound"
	_ "cyberservices.com/core/proxy/vless/outbound"
	_ "cyberservices.com/core/proxy/vmess/inbound"
	_ "cyberservices.com/core/proxy/vmess/outbound"

	// Transports
	_ "cyberservices.com/core/transport/internet/domainsocket"
	_ "cyberservices.com/core/transport/internet/http"
	_ "cyberservices.com/core/transport/internet/kcp"
	_ "cyberservices.com/core/transport/internet/quic"
	_ "cyberservices.com/core/transport/internet/tcp"
	_ "cyberservices.com/core/transport/internet/tls"
	_ "cyberservices.com/core/transport/internet/udp"
	_ "cyberservices.com/core/transport/internet/websocket"
	_ "cyberservices.com/core/transport/internet/xtls"

	// Transport headers
	_ "cyberservices.com/core/transport/internet/headers/http"
	_ "cyberservices.com/core/transport/internet/headers/noop"
	_ "cyberservices.com/core/transport/internet/headers/srtp"
	_ "cyberservices.com/core/transport/internet/headers/tls"
	_ "cyberservices.com/core/transport/internet/headers/utp"
	_ "cyberservices.com/core/transport/internet/headers/wechat"
	_ "cyberservices.com/core/transport/internet/headers/wireguard"

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	_ "cyberservices.com/core/main/json"
	// The following line loads JSON internally
	// _ "cyberservices.com/core/main/jsonem"

	// Load config from file or http(s)
	_ "cyberservices.com/core/main/confloader/external"
)
