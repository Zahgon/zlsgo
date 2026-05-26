package auth

import (
	"github.com/sohaha/zlsgo/znet"
)

const UserKey = "auth_user"

type (
	Accounts  map[string]string
	authPairs []authPair
	authPair  struct {
		value string
		user  string
	}
)

func (a authPairs) searchCredential(authValue string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func New(accounts Accounts) znet.Handler { _ = "STUB: not implemented"; return *new(znet.Handler) }

func BasicRealm(accounts Accounts, realm string) znet.Handler {
	_ = "STUB: not implemented"
	return *new(znet.Handler)
}

func processAccounts(accounts Accounts) (authPairs, error) {
	_ = "STUB: not implemented"
	return *new(authPairs), nil
}

func authorizationHeader(user, password string) string { _ = "STUB: not implemented"; return "" }
