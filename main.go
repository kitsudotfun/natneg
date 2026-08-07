package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/netip"
	"slices"

	. "github.com/kitsudotfun/natneg/defs"
)

var (
	conn *net.UDPConn

	handlers = map[byte]HandlerFunc{
		KeepAlive: hand(handleKeepAlive),
		Discover:  hand(handleDiscover),
		Join:      hand(handleJoin),
	}
)

func main() {
	var err error
	conn, err = net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP("51.81.22.70"), Port: 62426})
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1400)
	for {
		n, addr, err := conn.ReadFromUDPAddrPort(buf)
		if err != nil {
			log.Printf("%s: error: %s", addr, err)
			continue
		}
		if n < 1 {
			log.Printf("%s: error: %s", addr, err)
			continue
		}

		data := slices.Clone(buf[:n])

		log.Printf("%s: got packet: %x", addr, data)

		handler, exists := handlers[data[0]]
		if !exists {
			log.Printf("%s: error: no handler for packet type %x", addr, data[0])
			continue
		}

		resp, err := handler(conn, addr, data[1:])
		if err != nil {
			log.Printf("%s: error: %s", addr, err)
			continue
		}

		_, err = conn.WriteToUDPAddrPort(append([]byte{data[0]}, resp...), addr)
		if err != nil {
			log.Printf("%s: error: %s", addr, err)
			continue
		}
	}
}

type HandlerFunc func(conn *net.UDPConn, addr netip.AddrPort, data []byte) ([]byte, error)

func hand[reqT any, resT any](handler func(reqT, netip.AddrPort) (resT, error)) HandlerFunc {
	return func(conn *net.UDPConn, addr netip.AddrPort, data []byte) ([]byte, error) {
		var req reqT
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&req)
		if err != nil {
			return nil, err
		}

		res, err := handler(req, addr)
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(res)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}
}
