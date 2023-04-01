package client

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/coming-chat/go-sui/types"

	"github.com/coming-chat/go-sui/account"
	"github.com/stretchr/testify/require"
)

var (
	M1Mnemonic   = os.Getenv("WalletSdkTestM1")
	M1Address, _ = types.NewAddressFromHex("0x7e875ea78ee09f08d72e2676cf84e0f1c8ac61d94fa339cc8e37cace85bebc6e")
	Address      = M1Address
)

var (
	out, _ = exec.Command("whoami").Output()
	whoami = strings.TrimSpace(string(out))
)

func TestnetClient(t *testing.T) *Client {
	c, err := Dial(types.TestnetRpcUrl)
	require.NoError(t, err)
	return c
}

func DevnetClient(t *testing.T) *Client {
	c, err := Dial(types.DevNetRpcUrl)
	require.NoError(t, err)

	coins, err := c.GetCoins(context.TODO(), *Address, nil, nil, 1)
	require.NoError(t, err)
	if len(coins.Data) == 0 {
		_, err = FaucetFundAccount(Address.String(), DevNetFaucetUrl)
		require.NoError(t, err)
	}
	return c
}

func M1Account(t *testing.T) *account.Account {
	a, err := account.NewAccountWithMnemonic(M1Mnemonic)
	require.NoError(t, err)
	return a
}

func Signer(t *testing.T) *account.Account {
	return M1Account(t)
}
