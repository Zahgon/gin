package gin

const AuthUserKey = "user"

const AuthProxyUserKey = "proxy_user"

type Accounts map[string]string

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func (a authPairs) searchCredential(authValue string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func BasicAuth(accounts Accounts) HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func processAccounts(accounts Accounts) authPairs {
	_ = "STUB: not implemented"
	return *new(authPairs)
}

func authorizationHeader(user, password string) string { _ = "STUB: not implemented"; return "" }

func BasicAuthForProxy(accounts Accounts, realm string) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}
