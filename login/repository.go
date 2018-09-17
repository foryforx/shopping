package login

import (
	"context"

	"github.com/karuppaiah/shopping/model"
)

type ERepository interface {
	//Fetch(ctx context.Context) ([]*model.Login, error)
	Store(ctx context.Context, a *model.Login) (int64, error)
	Delete(ctx context.Context, username string) (bool, error)
	FetchLoginwithUsername(ctx context.Context, query string) ([]*model.Login, error)
}
