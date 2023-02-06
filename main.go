package main

import (
	"encoding/hex"
	"flag"
	"mnemonic/hdwallet"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogather/com/log"
	"github.com/eoscanada/eos-go/btcsuite/btcd/btcec"
	"github.com/eoscanada/eos-go/btcsuite/btcutil"
)

func main()  {
	mnemonic := flag.String("mnemonic", "", "mnemonic")
	flag.Parse()
	log.Println("mnemonic:", *mnemonic)
	wallet, err := hdwallet.GetMyKeyWallet(*mnemonic)
	if err != nil {
		log.Println("mnemonic error:", err.Error())
		return
	}
	ethPrivate := hexutil.Encode(wallet.Key[1:])
	log.Println("ETH private:", hexutil.Encode(wallet.Key[1:]))
	eosPrivate, err := ConvertToWif(ethPrivate[2:])
	if err != nil {
		log.Println("ConvertToWif to eos error:", err.Error())
		return
	}
	log.Println("EOS private:", eosPrivate.String())
}

func ConvertToWif(privateKeyHex string) (wifFromETH *btcutil.WIF, err error) {
	privateKeyByte, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, err
	}
	privateKeyForBTC, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyByte)
	wifFromETH, err = btcutil.NewWIF(privateKeyForBTC, 0x80, false)
	if err != nil {
		return nil, err
	}
	return wifFromETH, err
}