package main

import (
	"net/netip"

	. "github.com/kitsudotfun/kyuubi/api/defs"
	. "github.com/kitsudotfun/natneg/defs"

	"github.com/golang-jwt/jwt/v5"
)

func handleDiscover(req DiscoverRequest, addr netip.AddrPort) (DiscoverResponse, error) {
	var claims SessionClaims
	token, err := jwt.ParseWithClaims(req.Token, &claims, func(t *jwt.Token) (any, error) { return MustGetJwtKey("session"), nil })
	if err != nil {
		return DiscoverResponse{}, err
	}
	if !token.Valid {
		return DiscoverResponse{}, ErrInvalidToken
	}

	ret, err := jwt.NewWithClaims(JwtMethod, NatNegClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"natneg_attest"},
			ExpiresAt: claims.ExpiresAt,
		},
		Session: claims.Session,
		Addr:    addr,
	}).SignedString(MustGetJwtKey("natneg"))
	if err != nil {
		return DiscoverResponse{}, err
	}

	return DiscoverResponse{
		Token: ret,
	}, nil
}
