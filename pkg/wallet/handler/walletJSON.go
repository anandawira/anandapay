package handler

import "github.com/anandawira/anandapay/domain"

type BalanceResponseData struct {
	WalletID string `json:"wallet_id"`
	Balance  uint64 `json:"balance"`
}

type TopupRequestData struct {
	Amount uint32 `form:"amount" binding:"required"`
}

type TopupResponseData struct {
	domain.Transaction
}
