package lib

import (
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
	tls "github.com/refraction-networking/utls"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

//tinyjson:json
type ShadowsocksConfig struct {
	dummy string
}

// MarshalJSON supports json.Marshaler interface
func (v ShadowsocksConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	// TODO
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ShadowsocksConfig) MarshalTinyJSON(w *jwriter.Writer) {
	// TODO
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShadowsocksConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	// TODO
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ShadowsocksConfig) UnmarshalTinyJSON(l *jlexer.Lexer) {
	// TODO
}

//tinyjson:json
type Configurables struct {
	TLSConfig         *tls.Config        `json:"tls_config"`
	ShadowsocksConfig *ShadowsocksConfig `json:"shadowsocks_config"`
}

// MarshalJSON supports json.Marshaler interface
func (v Configurables) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	// TODO
	return w.Buffer.BuildBytes(), w.Error
}

func (c *Configurables) GetTLSConfig() *tls.Config {
	return c.TLSConfig
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Configurables) MarshalTinyJSON(w *jwriter.Writer) {
	// TODO
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Configurables) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	// TODO
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Configurables) UnmarshalTinyJSON(l *jlexer.Lexer) {
	// TODO
}
