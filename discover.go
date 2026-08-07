package main

import (
	"net/netip"

	. "github.com/kitsudotfun/kyuubi/api/defs"
	. "github.com/kitsudotfun/natneg/defs"

	"github.com/golang-jwt/jwt/v5"
)

func handleDiscover(req DiscoverRequest, addr netip.AddrPort) (DiscoverResponse, error) {
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(req.Token, &claims, func(t *jwt.Token) (any, error) { return MustGetJwtKey("session"), nil })
	if err != nil {
		return DiscoverResponse{}, err
	}
	if !token.Valid {
		return DiscoverResponse{}, ErrInvalidToken
	}

	id, err := token.Claims.GetSubject()
	if err != nil {
		return DiscoverResponse{}, err
	}

	ret, err := jwt.NewWithClaims(JwtMethod, NatNegClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			Audience:  jwt.ClaimStrings{"natneg_attest"},
			ExpiresAt: claims.ExpiresAt,
		},
		Addr: addr,
	}).SignedString(MustGetJwtKey("natneg"))
	if err != nil {
		return DiscoverResponse{}, err
	}

	return DiscoverResponse{
		Token: ret,
	}, nil
}
