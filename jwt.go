package main

import (
	"context"
	"io"
	"net/http"
	"os"

	. "github.com/kitsudotfun/kyuubi/api/defs"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/kv"
)

var client = cloudflare.NewClient()

func MustGetJwtKey(id string) []byte {
	r, err := client.KV.Namespaces.Values.Get(context.TODO(), JwtKeyNamespace, id, kv.NamespaceValueGetParams{
		AccountID: cloudflare.String(os.Getenv("CLOUDFLARE_ACCOUNT_ID")),
	})
	if err != nil {
		panic(err)
	}
	if r.StatusCode != http.StatusOK {
		panic("non-ok status")
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	return b
}
