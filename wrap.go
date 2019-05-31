package ssh

import (
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
