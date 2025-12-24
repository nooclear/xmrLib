package xmrLib

import (
	"fmt"
	"net/http"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

// Wallet represents a JSON-RPC client destination for interacting with wallet-related functionalities.
type Wallet struct {
	jrpcLib.Destination
}
type WalletResponse struct {
	jrpcLib.JRPCResult
}

// NewWallet creates and returns a new Wallet instance configured with the specified HTTP client and connection details.
func NewWallet(client *http.Client, method, protocol, ip string, port int, path string) *Wallet {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:NewWallet", fmt.Sprintf("Creating New Client:\tClient: %v Method: %s Protocol: %s IP: %s Port: %d Path: %s", client, method, protocol, ip, port, path))
	}
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
