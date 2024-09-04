package domain

import "time"

type Bet struct {
	Date             time.Time `json:"date"`
	BetAmount        float64   `json:"bet_amount"`
	PossibleWin      float64   `json:"possible_win"`
	Outcome          string    `json:"outcome"`
	BettingCompanyID int64     `json:"betting_company_id"`
	CashoutAmount    float64   `json:"cashout_amount"`
}
