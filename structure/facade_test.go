package structure

import "testing"

func TestFacade(t *testing.T) {
	face := NewWalletFacade("wangxiaoba",7879)
	if err := face.AddMoneyToWallet("wangxiaoba", 7879, 10);err !=nil {
		t.Error(err)
	}
	if err := face.DeductMoneyToWallet("wangxiaoba", 7879, 20);err !=nil {
		t.Error(err)
	}
	if err := face.DeductMoneyToWallet("wangxiaoba", 7879, 5);err !=nil {
		t.Error(err)
	}
}