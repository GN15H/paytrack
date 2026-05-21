package accounts

type Account struct {
	ID        string  `db:"id" json:"id"`
	UserID    string  `db:"user_id" json:"user_id"`
	Name      string  `db:"name" json:"name"`
	Type      string  `db:"type" json:"type"`
	Balance   float64 `db:"balance" json:"balance"`
	Currency  string  `db:"currency" json:"currency"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}

type CreateAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=checking savings cash"`
	Currency string `json:"currency" binding:"required"`
}
