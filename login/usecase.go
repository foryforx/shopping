package login

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type EUsecase interface {
	Fetch(ctx context.Context, user string) ([]*model.Login, error)
	Store(ctx context.Context, a *model.Login) (int64, error)
	Delete(ctx context.Context, user string) (bool, error)
}
