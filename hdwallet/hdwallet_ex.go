package hdwallet

import (
	"github.com/tyler-smith/go-bip39"
	"log"
	"mnemonic/bip39Helper"
)

func NewSeedFromMnemonicWithPassword(mnemonic, password string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}

// NewSeedFromMnemonic returns a BIP-39 seed based on a BIP-39 mnemonic.
func NewSeedFromMnemonic(mnemonic string) ([]byte, error) {
	return NewSeedFromMnemonicWithPassword(mnemonic, "")
}

func GetMyKeyWallet(mnemonic string) (*HDWallet, error) {
	englishMnemonic, err := bip39Helper.GetEnglishMnemonic(mnemonic)
	if err != nil {
		log.Println("in getMyKeyPub error:", err.Error())
		return nil, err
	}
	return NewWalletFromMnemonicForMyKey(englishMnemonic)
}

func NewWalletFromMnemonicForMyKey(mnemonic string) (*HDWallet, error) {
	seed, err := NewSeedFromMnemonic(mnemonic)
	if err != nil {
		log.Println("in newWalletFromMnemonicForMyKey error:", err.Error())
		return nil, err
	}
	return MasterKey(seed), nil
}
