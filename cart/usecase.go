package cart

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type EUsecase interface {
	Fetch(ctx context.Context, user string) ([]*model.Cart, error)
	Store(ctx context.Context, a *model.Cart) (int64, error)
	Delete(ctx context.Context, id int) (bool, error)
}
