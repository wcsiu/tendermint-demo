package demo

import "github.com/tendermint/tendermint/crypto/secp256k1"

func PrivateKey(key secp256k1.PrivKey) secp256k1.PrivKey {
	return key
}
