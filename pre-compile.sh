source variables.sh

echo "// generated
package main

const mode = \"$CompileMode\"

func wgConfig() string {
	return \`[Interface]
Address = 10.7.0.2/24
DNS = 10.8.0.1, 8.8.8.8
PrivateKey = $InterfacePrivateKey
[Peer]
PublicKey = $PeerPublicKey
PresharedKey = $PeerPresharedKey
AllowedIPs = $LfVpnIp/32
Endpoint = $JumpServerIp:$JumpServerPort
PersistentKeepalive = 25\`
}
" > wg_config.go

exit 0
