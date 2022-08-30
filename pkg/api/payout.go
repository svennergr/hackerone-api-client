package api

import "time"

type Payout struct {
	Amount         float32       `json:"amount"`
	PaidOutAt      time.Time `json:"paid_out_at"`
	Reference      string    `json:"reference"`
	PayoutProvider string    `json:"payout_provider"`
	Status         string    `json:"status"`
}
