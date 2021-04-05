package diffiehellman

import (
	"crypto/rand"
	"log"
	"math/big"
)

var (
	// big0 is used at the beginning
	// of each function call so that
	// the passed-in big.Int is not
	// modified
	big0 = big.NewInt(0)
	big2 = big.NewInt(2)
)

func PrivateKey(p *big.Int) *big.Int {
	// Creates a key > 1 && < p
	// 3 is subtracted because  the min is 2, and the
	// max is p - 1, so 3 in total
	a, err := rand.Int(rand.Reader, big0.Sub(p, big2))
	if err != nil {
		log.Fatal(err)
	}

	// add 2 since it must be > 1
	return a.Add(a, big2)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	// A = g**a mod p
	return big0.Exp(big.NewInt(g), private, p)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	// s = B**a mod p
	return big0.Exp(public2, private1, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}
