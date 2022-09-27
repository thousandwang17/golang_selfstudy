package main

type walletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) (*walletFacade, error) {

	w := &walletFacade{
		account:      NewAccount(accountId),
		wallet:       NewWallet(),
		securityCode: NewSecurityCode(code),
		notification: &Notification{},
		ledger:       NewLedger(),
	}

	if err := w.account.CheckAcount(accountId); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *walletFacade) addMoneyToWallet(accountId string, code int, amount int) error {

	if err := w.account.CheckAcount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(code); err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	go w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}

func (w *walletFacade) deductMoneyWallet(accountId string, code int, amount int) error {
	if err := w.account.CheckAcount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(code); err != nil {
		return err
	}

	w.wallet.debitBalance(amount)
	go w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}
