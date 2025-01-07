package repository

import (
	"context"
	"database/sql"

	"github.com/acronix0/XML-Parser-Golang/internal/repository/category"
	categoryModel "github.com/acronix0/XML-Parser-Golang/internal/repository/category/model"
	"github.com/acronix0/XML-Parser-Golang/internal/repository/product"
	productModel "github.com/acronix0/XML-Parser-Golang/internal/repository/product/model"
)

type Category interface {
	UpdateOrCreate(ctx context.Context, categories []categoryModel.Category) error
}

type Product interface {
	UpdateOrCreate(ctx context.Context, products []productModel.Product) error
}

type RepositoryManager interface{
	Category() Category
	Product() Product
}

type repositories struct {
	db *sql.DB
	category Category
	product Product
}

func NewRepositoryManager(db *sql.DB) *repositories{
	return &repositories{
		db: db,
	}
}

func (r *repositories) Category() Category{
	if r.category == nil {
		r.category = category.NewCategoryRepository(r.db)
	}
	return r.category
}

func (r *repositories) Product() Product{
	if r.product == nil {
		r.product = product.NewProductRepository(r.db)
	}
	return r.product
}

