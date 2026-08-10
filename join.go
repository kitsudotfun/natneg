package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/netip"
	"slices"

	. "github.com/kitsudotfun/kyuubi/api/defs"
	. "github.com/kitsudotfun/natneg/defs"

	"github.com/golang-jwt/jwt/v5"
)

func handleJoin(req JoinRequest, addr netip.AddrPort) (JoinResponse, error) {
	var claims ServerJoinClaims
	token, err := jwt.ParseWithClaims(req.Token, &claims, func(t *jwt.Token) (any, error) { return MustGetJwtKey("natneg"), nil })
	if err != nil {
		return JoinResponse{}, err
	}
	if !token.Valid || !slices.Contains(claims.Audience, "natneg_join") {
		return JoinResponse{}, ErrInvalidToken
	}

	var sid SessionID
	sid.FromString(claims.Subject)

	var buf bytes.Buffer
	buf.WriteString(NatnegMagic)
	buf.WriteByte(JoinNotify)
	err = json.NewEncoder(&buf).Encode(JoinNotifyResponse{
		ClientID:   sid,
		ClientAddr: addr,
	})
	if err != nil {
		return JoinResponse{}, err
	}

	log.Printf("%s: sent packet: %x", claims.ServerAddr, buf.Bytes())

	_, err = conn.WriteToUDPAddrPort(buf.Bytes(), claims.ServerAddr)
	if err != nil {
		return JoinResponse{}, err
	}

	return JoinResponse{
		ServerID:   claims.ServerID,
		ServerAddr: claims.ServerAddr,
	}, nil
}
