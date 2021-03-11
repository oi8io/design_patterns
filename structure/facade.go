package structure

import (
	"errors"
	"fmt"
)

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{name: name}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect ")
	}
	fmt.Println("Account Verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
}
func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("余额不足")
	}
	w.balance -= amount
	fmt.Printf("pay %d balance %d\n", amount, w.balance)
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{code: code}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect ")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func NewWalletFacade(accountName string, securityCode int) *walletFacade {
	return &walletFacade{
		account:      newAccount(accountName),
		wallet:       newWallet(),
		securityCode: newSecurityCode(securityCode),
		notification: &notification{},
		ledger:       &ledger{},
	}
}

func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) DeductMoneyToWallet(accountID string, securityCode int, amount int) error {
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}
	if err := w.wallet.debitBalance(amount); err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
