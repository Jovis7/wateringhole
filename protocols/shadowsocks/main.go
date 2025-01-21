package main

import v1 "github.com/refraction-networking/watm/tinygo/v1"

func init() {
	v1.BuildDialerWithDialingTransport(&ShadowsocksTransport{})
	// v1.BuildRelayWithWrappingTransport(&ShadowsocksTransport{}, v1.RelayWrapRemote)
}

func main() {}
