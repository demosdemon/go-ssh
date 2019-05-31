package ssh

import (
	"crypto/rand"
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

func generateSigner() (Signer, error) {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	return ssh.NewSignerFromKey(key)
}
