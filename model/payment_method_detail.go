package model

type PaymentMethodDetail struct {
	Type         *string         `json:"type,omitempty"`
	MandateID    *string         `json:"mandate_id,omitempty"`
	Card         *CardDetail     `json:"card,omitempty"`
	AchDebit     *AchDebitDetail `json:"ach_debit,omitempty"`
	Upi          *Upi            `json:"upi,omitempty"`
	NetBanking   *NetBanking     `json:"net_banking,omitempty"`
	BankTransfer *BankTransfer   `json:"bank_transfer,omitempty"`
}
