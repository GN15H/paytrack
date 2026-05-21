package summary

import "github.com/jmoiron/sqlx"

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetSummary(userID, from, to string) (Summary, error) {
	var summary Summary
	err := s.db.QueryRowx(
		`SELECT
			COALESCE(SUM(CASE WHEN t.type = 'income' THEN t.amount ELSE 0 END), 0) AS total_income,
			COALESCE(SUM(CASE WHEN t.type = 'expense' THEN t.amount ELSE 0 END), 0) AS total_expense,
			COALESCE(SUM(CASE WHEN t.type = 'income' THEN t.amount
			               WHEN t.type = 'expense' THEN -t.amount
			               ELSE 0 END), 0) AS balance
		FROM transactions t
		JOIN accounts a ON a.id = t.account_id
		WHERE a.user_id = $1
		  AND t.date BETWEEN $2 AND $3`,
		userID, from, to,
	).StructScan(&summary)
	return summary, err
}

func (s *Service) GetByCategory(userID, from, to string) ([]CategorySummary, error) {
	var results []CategorySummary
	err := s.db.Select(&results,
		`SELECT
			t.category_id,
			c.name AS category_name,
			t.type,
			SUM(ABS(t.amount)) AS total
		FROM transactions t
		JOIN accounts a ON a.id = t.account_id
		LEFT JOIN categories c ON c.id = t.category_id
		WHERE a.user_id = $1
		  AND t.date BETWEEN $2 AND $3
		  AND t.type IN ('income', 'expense')
		GROUP BY t.category_id, c.name, t.type
		ORDER BY total DESC`,
		userID, from, to,
	)
	return results, err
}
