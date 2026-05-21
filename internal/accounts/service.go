package accounts

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(userID string, req CreateAccountRequest) (Account, error) {
	var account Account
	err := s.db.QueryRowx(
		`INSERT INTO accounts (user_id, name, type, currency)
		 VALUES ($1, $2, $3, $4)
		 RETURNING *`,
		userID, req.Name, req.Type, req.Currency,
	).StructScan(&account)
	return account, err
}

func (s *Service) FindAll(userID string) ([]Account, error) {
	var accounts []Account
	err := s.db.Select(&accounts,
		`SELECT * FROM accounts WHERE user_id = $1 ORDER BY created_at DESC`,
		userID,
	)
	return accounts, err
}

func (s *Service) FindByID(userID, accountID string) (Account, error) {
	var account Account
	err := s.db.Get(&account,
		`SELECT * FROM accounts WHERE id = $1 AND user_id = $2`,
		accountID, userID,
	)
	return account, err
}

func (s *Service) Delete(userID, accountID string) error {
	_, err := s.db.Exec(
		`DELETE FROM accounts WHERE id = $1 AND user_id = $2`,
		accountID, userID,
	)
	return err
}
