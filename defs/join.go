package defs

import (
	"net/netip"

	. "github.com/kitsudotfun/kyuubi/api/defs"
)

type JoinRequest struct {
	Token string `json:"token"`
}
type JoinResponse struct {
	ServerID   PeerID         `json:"server_id"`
	ServerAddr netip.AddrPort `json:"server_addr"`
}

type JoinNotifyResponse struct {
	ClientID   PeerID         `json:"client_id"`
	ClientAddr netip.AddrPort `json:"client_addr"`
}
