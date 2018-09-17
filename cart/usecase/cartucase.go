package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/karuppaiah/shopping/cart"
	"github.com/karuppaiah/shopping/model"
)

// eUsecase having link to Repository
type eUsecase struct {
	eRepos         cart.ERepository
	contextTimeout time.Duration
}

// NewEUsecase :To Create new usecase for calling from API layer
func NewEUsecase(a cart.ERepository, timeout time.Duration) cart.EUsecase {
	return &eUsecase{
		eRepos:         a,
		contextTimeout: timeout,
	}
}

//Fetch the cart items for the particular user
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

// Get total value after discount for the list of cart items
func (a *eUsecase) GetTotalCartValue(m []*model.Cart) float64 {
	total := 0.0
	for i := 0; i < len(m); i++ {
		total = total + (float64(m[i].Items) * m[i].Price) - m[i].Dprice
	}
	return total
}

// Add a new cart items
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

// Delete the cart items
func (a *eUsecase) Delete(c context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.eRepos.Delete(ctx, id)
}

// Refresh the cart to update all Discount price for the particular user
func (a *eUsecase) refresh(c context.Context, user string) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	// Get cart items in map format
	mCart, err := a.eRepos.ConvertCartDetailsAsMap(ctx, user)
	fmt.Println("Refresh:", mCart)
	if err != nil {
		return err
	}
	// Get the list of promotions applicable for this cart item
	lPromo, err := a.eRepos.FetchPromotionDetailsForCart(ctx, user)
	if err != nil {
		return err
	}

	// Set previous applied promotion to 0[since some promotion might be deleted from system]
	for _, v := range mCart {
		if v.Dprice != 0 {
			v.Dprice = 0
			_, err := a.eRepos.Update(ctx, v)
			if err != nil {
				return err
			}
		}
	}
	//For each promotion in desc priority, calculate DPrice for cart
	for i := 0; i < len(lPromo); i++ {
		fmt.Println(i, ".Source product id present:", mCart[lPromo[i].Sprodid])
		sCartItem := mCart[lPromo[i].Sprodid]
		dCartItem := mCart[lPromo[i].Dprodid]
		fmt.Println("Source cart item:", sCartItem)
		fmt.Println("Destination Cart items:", dCartItem)
		// If min source product item present and dest. cart item present?
		if sCartItem.Items >= lPromo[i].Sminqty && dCartItem != nil {
			fmt.Println("Promotion applicable")
			noOfItemDiscApplied := 0
			// If dest. cart items is more than min qualified promotion qty
			if dCartItem.Items > lPromo[i].Dminqty {
				noOfItemDiscApplied = dCartItem.Items - lPromo[i].Dminqty
				// If discount is percentage
				if lPromo[i].Disctype == "P" {
					dCartItem.Dprice = ((lPromo[i].Discount / 100) * dCartItem.Price) * float64(noOfItemDiscApplied)
				} else if lPromo[i].Disctype == "F" {
					// If discount is fixed discount
					dCartItem.Dprice = (lPromo[i].Discount) * float64(noOfItemDiscApplied)
				}
				fmt.Println("Final cart item:", dCartItem)
				// finally update cart item
				_, err := a.eRepos.Update(ctx, dCartItem)
				if err != nil {
					return err
				}
			}

		}

	}
	return nil
}
