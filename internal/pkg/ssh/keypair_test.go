package ssh_test

import (
	"bwizard/internal/pkg/ssh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestCreateSSHKeyPair(t *testing.T) {
	dir := t.TempDir()

	err := ssh.CreateSSHKeyPair(dir)
	require.NoError(t, err)

	privateKey, err := os.ReadFile(path.Join(dir, "id_ed25519"))
	require.NoError(t, err)

	publicKey, err := os.ReadFile(path.Join(dir, "id_ed25519.pub"))
	require.NoError(t, err)

	assert.True(t, len(privateKey) > 0)
	assert.True(t, len(publicKey) > 0)
}

func TestCreateSSHKeyPairIfNotExists(t *testing.T) {
	dir := t.TempDir()

	err := ssh.CreateSSHKeyPair(dir)
	require.NoError(t, err)

	oldPrivateKey, err := os.ReadFile(path.Join(dir, "id_ed25519"))
	require.NoError(t, err)

	oldPublicKey, err := os.ReadFile(path.Join(dir, "id_ed25519.pub"))
	require.NoError(t, err)

	err = ssh.CreateSSHKeyPairIfNotExists(dir)
	require.NoError(t, err)

	newPrivateKey, err := os.ReadFile(path.Join(dir, "id_ed25519"))
	require.NoError(t, err)

	newPublicKey, err := os.ReadFile(path.Join(dir, "id_ed25519.pub"))
	require.NoError(t, err)

	assert.Equal(t, oldPrivateKey, newPrivateKey)
	assert.Equal(t, oldPublicKey, newPublicKey)
}
