package promotion

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type ERepository interface {
	Fetch(ctx context.Context) ([]*model.Promotion, error)
	Store(ctx context.Context, a *model.Promotion) (int64, error)
	Delete(ctx context.Context, id int) (bool, error)
	FetchPromotionwithQuery(ctx context.Context, query string) ([]*model.Promotion, error)
}
