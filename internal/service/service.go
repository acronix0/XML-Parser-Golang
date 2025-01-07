package service

import "context"

type Parser interface {
	Parse(ctx context.Context, filePath string) error
}