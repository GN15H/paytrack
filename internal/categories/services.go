package categories

import "github.com/jmoiron/sqlx"

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(userID string, req CreateCategoryRequest) (Category, error) {
	var category Category
	err := s.db.QueryRowx(
		`INSERT INTO categories (user_id, name, type)
		 VALUES ($1, $2, $3)
		 RETURNING *`,
		userID, req.Name, req.Type,
	).StructScan(&category)
	return category, err
}

func (s *Service) FindAll(userID string) ([]Category, error) {
	var categories []Category
	err := s.db.Select(&categories,
		`SELECT * FROM categories
		 WHERE user_id = $1 OR user_id = '00000000-0000-0000-0000-000000000000'
		 ORDER BY name ASC`,
		userID,
	)
	return categories, err
}
