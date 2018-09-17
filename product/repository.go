package product

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type ProductRepository interface {
	Fetch(ctx context.Context) ([]*model.Product, error)
	Store(ctx context.Context, a *model.Product) (int64, error)
	Delete(ctx context.Context, id int) (bool, error)
	FetchProductWithQuery(ctx context.Context, query string) ([]*model.Product, error)
}
