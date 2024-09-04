-- name: CreateBet :one
INSERT INTO bets (
    user_id, 
    betting_company_id,
    bet_amount,
    outcome,
    possible_win,
    cashout_amount,
    date
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;
