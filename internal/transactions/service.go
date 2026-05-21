package transactions

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(userID, accountID string, req CreateTransactionRequest) (Transaction, error) {
	var count int
	err := s.db.Get(&count,
		`SELECT COUNT(*) FROM accounts WHERE id = $1 AND user_id = $2`,
		accountID, userID,
	)
	if err != nil || count == 0 {
		return Transaction{}, errors.New("account not found")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return Transaction{}, err
	}
	defer tx.Rollback()

	var transaction Transaction
	err = tx.QueryRowx(
		`INSERT INTO transactions (account_id, category_id, amount, type, description, date)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING *`,
		accountID, req.CategoryID, req.Amount, req.Type, req.Description, req.Date,
	).StructScan(&transaction)
	if err != nil {
		return Transaction{}, err
	}

	delta := req.Amount
	if req.Type == "expense" {
		delta = -req.Amount
	}

	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`,
		delta, accountID,
	)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, tx.Commit()
}

func (s *Service) Transfer(userID string, req TransferRequest) error {
	var count int
	err := s.db.Get(&count,
		`SELECT COUNT(*) FROM accounts WHERE id IN ($1, $2) AND user_id = $3`,
		req.FromAccountID, req.ToAccountID, userID,
	)
	if err != nil || count < 2 {
		return errors.New("one or both accounts not found")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	transferType := "transfer"

	_, err = tx.Exec(
		`INSERT INTO transactions (account_id, amount, type, description, date)
		 VALUES ($1, $2, $3, $4, $5)`,
		req.FromAccountID, -req.Amount, transferType, req.Description, req.Date,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`INSERT INTO transactions (account_id, amount, type, description, date)
		 VALUES ($1, $2, $3, $4, $5)`,
		req.ToAccountID, req.Amount, transferType, req.Description, req.Date,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance - $1 WHERE id = $2`,
		req.Amount, req.FromAccountID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`,
		req.Amount, req.ToAccountID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) FindAll(userID, accountID string) ([]Transaction, error) {
	var count int
	err := s.db.Get(&count,
		`SELECT COUNT(*) FROM accounts WHERE id = $1 AND user_id = $2`,
		accountID, userID,
	)
	if err != nil || count == 0 {
		return nil, errors.New("account not found")
	}

	var transactions []Transaction
	err = s.db.Select(&transactions,
		`SELECT * FROM transactions
		 WHERE account_id = $1
		 ORDER BY date DESC, created_at DESC`,
		accountID,
	)
	return transactions, err
}

func (s *Service) Delete(userID, accountID, transactionID string) error {
	var transaction Transaction
	err := s.db.Get(&transaction,
		`SELECT t.* FROM transactions t
		 JOIN accounts a ON a.id = t.account_id
		 WHERE t.id = $1 AND t.account_id = $2 AND a.user_id = $3`,
		transactionID, accountID, userID,
	)
	if err != nil {
		return errors.New("transaction not found")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`DELETE FROM transactions WHERE id = $1`, transactionID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance - $1 WHERE id = $2`,
		transaction.Amount, accountID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
