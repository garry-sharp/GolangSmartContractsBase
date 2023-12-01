package config

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

var config *Config
var privateKey *ecdsa.PrivateKey

type CoreConfig struct {
	EthProvider string
	ChainID     int64
}

// Extend this config
type Config struct {
	CoreConfig
}

func GetConfig() *Config {
	return config
}

// TODO implement this function
func getPrivateKey() (*ecdsa.PrivateKey, error) {
	var err error = nil
	if privateKey == nil {
		privateKey, err = crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") // Address 1 of the Test Test Test... mnemonic
	}
	return privateKey, err
}

func GetAuth() (*bind.TransactOpts, error) {
	key, err := getPrivateKey()
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(key, big.NewInt(config.ChainID))
}
