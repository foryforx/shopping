package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/karuppaiah/shopping/model"
	"github.com/karuppaiah/shopping/promotion"
)

type eUsecase struct {
	eRepos         promotion.ERepository
	contextTimeout time.Duration
}

func NewEUsecase(a promotion.ERepository, timeout time.Duration) promotion.EUsecase {
	return &eUsecase{
		eRepos:         a,
		contextTimeout: timeout,
	}
}

// Fetch promotion item
func (a *eUsecase) Fetch(c context.Context) ([]*model.Promotion, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listent, err := a.eRepos.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return listent, nil
}

// Save promotion item
func (a *eUsecase) Store(ctx context.Context, m *model.Promotion) (int64, error) {
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

// Delete promotion item
func (a *eUsecase) Delete(c context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.eRepos.Delete(ctx, id)
}
