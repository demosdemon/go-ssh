package ssh

import (
	"crypto/rand"
	"crypto/rsa"
)

func generateSigner() (Signer, error) {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	return NewSignerFromSigner(key)
}
