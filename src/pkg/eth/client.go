package eth

import (
	"xxx/src/pkg/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient() (*ethclient.Client, error) {
	return ethclient.Dial(config.GetConfig().EthProvider)
}
