package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/karuppaiah/shopping/cart"
	"github.com/karuppaiah/shopping/model"
)

type eUsecase struct {
	eRepos         cart.ERepository
	contextTimeout time.Duration
}

func NewEUsecase(a cart.ERepository, timeout time.Duration) cart.EUsecase {
	return &eUsecase{
		eRepos:         a,
		contextTimeout: timeout,
	}
}

func (a *eUsecase) Fetch(c context.Context, user string) ([]*model.Cart, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listent, err := a.eRepos.Fetch(ctx, user)
	if err != nil {
		return nil, err
	}

	return listent, nil
}

func (a *eUsecase) Store(ctx context.Context, m *model.Cart) (int64, error) {
	fmt.Println("uc:", m)
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	id, err := a.eRepos.Store(ctx, m)
	if err != nil {
		return 0, err
	}
	fmt.Println("id:", id)
	return id, nil
}

func (a *eUsecase) Delete(c context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.eRepos.Delete(ctx, id)
}
