package defs

import (
	"net/netip"

	. "github.com/kitsudotfun/kyuubi/api/defs"
)

type JoinRequest struct {
	Token string `json:"token"`
}
type JoinResponse struct {
	ServerID   SessionID      `json:"server_id"`
	ServerAddr netip.AddrPort `json:"server_addr"`
}

type JoinNotifyResponse struct {
	ClientID   SessionID      `json:"client_id"`
	ClientAddr netip.AddrPort `json:"client_addr"`
}
