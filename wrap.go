package ssh

import (
	"crypto"
	"golang.org/x/crypto/ssh"
)

type PublicKey interface {
	ssh.PublicKey
}

type Permissions struct {
	*ssh.Permissions
}

type Signer interface {
	ssh.Signer
}

func ParseAuthorizedKey(in []byte) (out PublicKey, comment string, options []string, rest []byte, err error) {
	return ssh.ParseAuthorizedKey(in)
}

func ParsePublicKey(in []byte) (out PublicKey, err error) {
	return ssh.ParsePublicKey(in)
}

func NewSignerFromSigner(signer crypto.Signer) (Signer, error) {
	return ssh.NewSignerFromSigner(signer)
}

func NewSignerFromKey(key interface{}) (Signer, error) {
	return ssh.NewSignerFromKey(key)
}
