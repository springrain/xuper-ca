/*
 * Copyright (c) 2019. Baidu Inc. All Rights Reserved.
 */
package crypto

import (
	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/xuperchain/core/crypto/client/base"
	eccdefault "github.com/xuperchain/xuperchain/core/crypto/client/xchain"
)

// var CryptoTypeConfig = crypto_client.CryptoTypeDefault
func getInstance() interface{} {
	return &eccdefault.XchainCryptoClient{}
}

// GetCryptoClient get crypto client
func GetCryptoClient() base.CryptoClient {
	cryptoClient := getInstance().(base.CryptoClient)
	return cryptoClient
}

// Get HdCrypto Client
func GetHdCryptoClient() *eccdefault.XchainCryptoClient {
	return new(eccdefault.XchainCryptoClient)
}

func GetGMHdCryptoClient() *gm.GmCryptoClient {
	return new(gm.GmCryptoClient)
}
