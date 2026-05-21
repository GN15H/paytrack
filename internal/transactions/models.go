package transactions

type Transaction struct {
	ID          string  `db:"id" json:"id"`
	AccountID   string  `db:"account_id" json:"account_id"`
	CategoryID  *string `db:"category_id" json:"category_id"`
	Amount      float64 `db:"amount" json:"amount"`
	Type        string  `db:"type" json:"type"`
	Description *string `db:"description" json:"description"`
	Date        string  `db:"date" json:"date"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
}

type CreateTransactionRequest struct {
	CategoryID  *string `json:"category_id"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Type        string  `json:"type" binding:"required,oneof=income expense transfer"`
	Description *string `json:"description"`
	Date        string  `json:"date" binding:"required"`
}

type TransferRequest struct {
	FromAccountID string  `json:"from_account_id" binding:"required"`
	ToAccountID   string  `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Description   *string `json:"description"`
	Date          string  `json:"date" binding:"required"`
}
