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
	refErr := a.refresh(ctx, user)
	if refErr != nil {
		return nil, refErr
	}
	listent, err := a.eRepos.Fetch(ctx, user)

	if err != nil {
		return nil, err
	}

	return listent, nil
}
func (a *eUsecase) GetTotalCartValue(m []*model.Cart) float64 {
	total := 0.0
	for i := 0; i < len(m); i++ {
		total = total + (float64(m[i].Items) * m[i].Price) - m[i].Dprice
	}
	return total
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

func (a *eUsecase) refresh(c context.Context, user string) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	mCart, err := a.eRepos.ConvertCartDetailsAsMap(ctx, user)
	fmt.Println("Refresh:", mCart)
	if err != nil {
		return err
	}
	lPromo, err := a.eRepos.FetchPromotionDetailsForCart(ctx, user)
	if err != nil {
		return err
	}
	for i := 0; i < len(lPromo); i++ {
		fmt.Println(i, ".Source product id present:", mCart[lPromo[i].Sprodid])
		sCartItem := mCart[lPromo[i].Sprodid]
		dCartItem := mCart[lPromo[i].Dprodid]
		fmt.Println("Source cart item:", sCartItem)
		fmt.Println("Destination Cart items:", dCartItem)

		if sCartItem.Items >= lPromo[i].Sminqty && dCartItem != nil {
			fmt.Println("Promotion applicable")
			noOfItemDiscApplied := 0
			if dCartItem.Items > lPromo[i].Dminqty {
				noOfItemDiscApplied = dCartItem.Items - lPromo[i].Dminqty
				if lPromo[i].Disctype == "P" {
					dCartItem.Dprice = ((lPromo[i].Discount / 100) * dCartItem.Price) * float64(noOfItemDiscApplied)
				} else if lPromo[i].Disctype == "F" {
					dCartItem.Dprice = (lPromo[i].Discount) * float64(noOfItemDiscApplied)
				}
				fmt.Println("Final cart item:", dCartItem)
				_, err := a.eRepos.Update(ctx, dCartItem)
				if err != nil {
					return err
				}
			}

		}

	}
	return nil
}
