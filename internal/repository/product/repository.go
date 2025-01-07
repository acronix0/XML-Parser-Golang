package product

import "database/sql"

type ProductReository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductReository{
	return &ProductReository{
		db: db,
	}
}