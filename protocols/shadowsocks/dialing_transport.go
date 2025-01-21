package main

import (
	"shadowsocks-watm/lib"

	"github.com/CosmWasm/tinyjson"
	tls "github.com/refraction-networking/utls"
	v1 "github.com/refraction-networking/watm/tinygo/v1"
	v1net "github.com/refraction-networking/watm/tinygo/v1/net"
)

// type guard: ShadowsocksTransport must implement [v1.DialingTransport].
var _ v1.DialingTransport = (*ShadowsocksTransport)(nil)

type ShadowsocksTransport struct {
	tlsConfig *tls.Config
	dialer    func(network, address string) (v1net.Conn, error)
}

func (ssdt *ShadowsocksTransport) SetDialer(dialer func(network, address string) (v1net.Conn, error)) {
	ssdt.dialer = dialer
}

func (ssdt *ShadowsocksTransport) Dial(network, address string) (v1net.Conn, error) {
	conn, err := ssdt.dialer(network, address)
	if err != nil {
		return nil, err
	}

	return conn, conn.SetNonBlock(true) // must set non-block, otherwise will block on read and lose fairness
}

// type guard: ShadowsocksTransport must implement [v1.ConfigurableTransport].
var _ v1.ConfigurableTransport = (*ShadowsocksTransport)(nil)

func (ssdt *ShadowsocksTransport) Configure(config []byte) error {
	configurables := &lib.Configurables{}
	if err := tinyjson.Unmarshal(config, configurables); err != nil {
		return err
	}
	ssdt.tlsConfig = configurables.GetTLSConfig()
	// TODO
	return nil
}
