package ssh

import (
	"golang.org/x/crypto/ed25519"
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
	privateKey, publicKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dir, "id_ed25519"), privateKey, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dir, "id_ed25519.pub"), publicKey, 0600); err != nil {
		return err
	}

	return nil
}
