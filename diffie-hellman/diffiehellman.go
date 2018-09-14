package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PublicKey returns a public key.
func PublicKey(a *big.Int, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), a, p)
}

// SecretKey returns a secret key.
func SecretKey(a *big.Int, B *big.Int, p *big.Int) *big.Int {
	return new(big.Int).Exp(B, a, p)
}

// PrivateKey returns a private key.
func PrivateKey(p *big.Int) *big.Int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return new(big.Int).Add(big.NewInt(2), new(big.Int).Rand(r, new(big.Int).Sub(p, big.NewInt(2))))
}

// NewPair returns new pair keys.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	pk := PrivateKey(p)
	return pk, PublicKey(pk, p, g)
}
