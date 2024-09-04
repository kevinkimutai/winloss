CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE betting_company (
    company_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE bets (
    bet_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL ON DELETE CASCADE,
    betting_company_id BIGINT REFERENCES betting_company(company_id) NOT NULL ON DELETE CASCADE,
    bet_amount DECIMAL(10, 2) NOT NULL CHECK (bet_amount > 0),
    outcome VARCHAR(50) DEFAULT 'Pending' CHECK (outcome IN ('Pending', 'Win', 'Loss', 'Cashout')),
    possible_win DECIMAL(10, 2) NOT NULL,
    cashout_amount DECIMAL(10, 2),
    date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
