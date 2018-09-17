package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/karuppaiah/shopping/login"
	"github.com/karuppaiah/shopping/model"
)

type eUsecase struct {
	eRepos         login.ERepository
	contextTimeout time.Duration
}

func NewEUsecase(a login.ERepository, timeout time.Duration) login.EUsecase {
	return &eUsecase{
		eRepos:         a,
		contextTimeout: timeout,
	}
}

// Fetch login item
func (a *eUsecase) Fetch(c context.Context, user string) ([]*model.Login, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listEnt, err := a.eRepos.FetchLoginwithUsername(ctx, user)
	if err != nil {
		return nil, err
	}

	return listEnt, nil
}

// Save login item
func (a *eUsecase) Store(ctx context.Context, m *model.Login) (int64, error) {
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

// Delete login item
func (a *eUsecase) Delete(c context.Context, user string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.eRepos.Delete(ctx, user)
}
