package summary

type Summary struct {
	TotalIncome  float64 `db:"total_income" json:"total_income"`
	TotalExpense float64 `db:"total_expense" json:"total_expense"`
	Balance      float64 `db:"balance" json:"balance"`
}

type CategorySummary struct {
	CategoryID   *string `db:"category_id" json:"category_id"`
	CategoryName *string `db:"category_name" json:"category_name"`
	Type         string  `db:"type" json:"type"`
	Total        float64 `db:"total" json:"total"`
}
