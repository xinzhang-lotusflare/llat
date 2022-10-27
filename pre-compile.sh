source variables.sh

echo "// generated
package main

const mode = \"$CompileMode\"

func wgConfig(ipAddress string, privateKey string, publicKey string, presharedKey string) string {
	return \`[Interface]
Address = \` + ipAddress + \`
DNS = $LfDns, $SecondaryDns
PrivateKey = \` + privateKey + \`
[Peer]
PublicKey = \` + publicKey + \`
PresharedKey = \` + presharedKey + \`
AllowedIPs = $LfVpnIp/32
Endpoint = $JumpServerIp:$JumpServerPort
PersistentKeepalive = 25\`
}
" > wg_config.go

exit 0
