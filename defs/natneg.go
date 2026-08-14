package defs

const PeerMagic = "KTsu"

const (
	KeepAlive = iota
	Discover
	Join
	JoinNotify // server to client only
)
