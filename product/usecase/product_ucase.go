package usecase

import (
	"context"
	"time"

	"github.com/karuppaiah/shopping/model"
	"github.com/karuppaiah/shopping/product"
)

type productUsecase struct {
	productRepos   product.ProductRepository
	contextTimeout time.Duration
}

func NewProductUsecase(a product.ProductRepository, timeout time.Duration) product.ProductUsecase {
	return &productUsecase{
		productRepos:   a,
		contextTimeout: timeout,
	}
}

// Get Product
func (a *productUsecase) Fetch(c context.Context) ([]*model.Product, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listProduct, err := a.productRepos.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return listProduct, nil
}

// Save product
func (a *productUsecase) Store(ctx context.Context, m *model.Product) (int64, error) {

	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	id, err := a.productRepos.Store(ctx, m)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Delete product
func (a *productUsecase) Delete(c context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.productRepos.Delete(ctx, id)
}
