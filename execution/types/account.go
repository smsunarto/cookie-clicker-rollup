package types

type Account struct {
	// AccountHash is stored in StateDB so we don't have to recompute it
	// everytime we calculate state hash
	AccountHash string // Hash(Nonce, ClickCount) -> AccountHash
	Nonce       int
	ClickCount  int
}
