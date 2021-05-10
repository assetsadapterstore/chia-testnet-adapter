package chiatestnet


import (
	"github.com/assetsadapterstore/chia-adapter/chia"
	"github.com/blocktree/openwallet/v2/log"
)

const (
	Symbol = "TXCH"
)

type WalletManager struct {
	*chia.WalletManager
}


func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = chia.NewWalletManager()
	wm.Config = chia.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "CHIA-TESTNET"
}