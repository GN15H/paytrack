package categories

type Category struct {
	ID        string `db:"id" json:"id"`
	UserID    string `db:"user_id" json:"user_id"`
	Name      string `db:"name" json:"name"`
	Type      string `db:"type" json:"type"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required,oneof=income expense"`
}
