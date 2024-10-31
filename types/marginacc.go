package types

type MarginAccountEvent struct {
	MarginAccAddr string
	Owner         string
	AccNumber     string
}

func NewMarginAccountEvent(
	marginAccAddr string,
	owner string,
	accNumber string,
) MarginAccountEvent {
	return MarginAccountEvent{
		MarginAccAddr: marginAccAddr,
		Owner:         owner,
		AccNumber:     accNumber,
	}
}

type WithdrawEvent struct {
	MarginAccAddr string
	Recipient     string
	WithdrawCoin  string
}

func NewWithdrawEvent(
	marginAccAddr string,
	recipient string,
	withdrawCoin string,
) WithdrawEvent {
	return WithdrawEvent{
		MarginAccAddr: marginAccAddr,
		Recipient:     recipient,
		WithdrawCoin:  withdrawCoin,
	}
}
