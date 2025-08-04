package xmrLib

import (
	"github.com/nooclear/jrpcLib"
	"net/http"
)

// Wallet represents a JSON-RPC client destination for interacting with wallet-related functionalities.
type Wallet struct {
	jrpcLib.Destination
}

// NewWallet creates and returns a new Wallet instance configured with the specified HTTP client and connection details.
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
