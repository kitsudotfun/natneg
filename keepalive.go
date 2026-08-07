package main

import (
	"net/netip"

	. "github.com/kitsudotfun/natneg/defs"
)

func handleKeepAlive(_ KeepAliveRequest, _ netip.AddrPort) (KeepAliveResponse, error) {
	return KeepAliveResponse{}, nil
}
