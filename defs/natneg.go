package defs

const NatnegMagic = "KTnn"

const (
	KeepAlive = iota
	Discover
	Join
	JoinNotify // server to client only
)
