package ssh

import (
	"crypto/rand"
	"encoding/pem"
	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"os"
	"path"
)

// CreateSSHKeyPairIfNotExists creates a new ed25519 key pair inside dir only if no key exists
func CreateSSHKeyPairIfNotExists(dir string) error {
	if _, err := os.Stat(path.Join(dir, "id_ed25519")); os.IsNotExist(err) {
		return CreateSSHKeyPair(dir)
	}

	return nil
}

// CreateSSHKeyPair creates a new ed25519 key pair inside dir
func CreateSSHKeyPair(dir string) error {
	pubKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	publicKey, err := ssh.NewPublicKey(pubKey)
	if err != nil {
		return err
	}

	privateKeyBytes := edkey.MarshalED25519PrivateKey(privateKey)
	privateKeyPem := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := os.WriteFile(path.Join(dir, "id_ed25519"), pem.EncodeToMemory(privateKeyPem), 0600); err != nil {
		return err
	}

	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	if err := os.WriteFile(path.Join(dir, "id_ed25519.pub"), authorizedKey, 0644); err != nil {
		return err
	}

	return nil
}
