package cart

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type ERepository interface {
	Fetch(ctx context.Context, user string) ([]*model.Cart, error)
	Store(ctx context.Context, a *model.Cart) (int64, error)
	Delete(ctx context.Context, id int) (bool, error)
	Update(ctx context.Context, ar *model.Cart) (*model.Cart, error)
	FetchPromotionDetailsForCart(ctx context.Context, user string) ([]*model.Promotion, error)
	ConvertCartDetailsAsMap(ctx context.Context, user string) (map[int](*model.Cart), error)
}
