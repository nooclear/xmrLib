package xmrLib

import (
	"github.com/nooclear/jrpcLib"
	"net/http"
)

type Wallet struct {
	jrpcLib.Destination
}

func NewWallet(client *http.Client, method, protocol, ip string, port int, path string) *Wallet {
	return &Wallet{
		Destination: jrpcLib.Destination{
			Client:   client,
			Method:   method,
			Protocol: protocol,
			IP:       ip,
			Port:     port,
			Path:     path,
		},
	}
}
