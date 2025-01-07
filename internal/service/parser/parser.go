package parser

import (
	"log/slog"

	"github.com/acronix0/XML-Parser-Golang/internal/repository"
)

type ParserService struct {
	log *slog.Logger
	categoryRepo repository.Category
	productRepo repository.Product
	firstFilePath string
	secondFilePath string
}

func NewParserService(
	log *slog.Logger,
	categoryRepo repository.Category,
	productRepo repository.Product,
	firstFilePath string,
	secondFilePath string,
) *ParserService{
	return &ParserService{
		log: log,
		categoryRepo: categoryRepo,
		productRepo: productRepo,
		firstFilePath: firstFilePath,
		secondFilePath: secondFilePath,
	}
}