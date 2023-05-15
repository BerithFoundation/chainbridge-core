package evmclient

import (
	"testing"

	"github.com/ChainSafe/chainbridge-core/crypto/secp256k1"
	"github.com/stretchr/testify/require"
)

func TestGetLatestBlock(t *testing.T) {
	kp, err := secp256k1.NewKeypairFromString("000000000000000000000000000000000000000000000000000000000000007B")
	require.NoError(t, err)

	c, err := NewEVMClient("https://api.baobab.klaytn.net:8651", kp)
	require.NoError(t, err)

	b, err := c.LatestBlock()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.Greater(t, b.Int64(), int64(0))

	t.Log(b.Int64())
}
