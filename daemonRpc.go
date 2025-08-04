package xmrLib

import (
	"github.com/nooclear/jrpcLib"
	"net/http"
)

func NewDaemon(client *http.Client, method, protocol, ip string, port int, path string) *jrpcLib.Destination {
	return &jrpcLib.Destination{
		Client:   client,
		Method:   method,
		Protocol: protocol,
		IP:       ip,
		Port:     port,
		Path:     path,
	}
}
