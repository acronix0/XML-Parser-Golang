package category

import "database/sql"

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepo{
	return &CategoryRepo{
		db: db,
	}
}