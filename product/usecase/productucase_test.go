package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/karuppaiah/shopping/model"
	"github.com/karuppaiah/shopping/product/mocks"
	"github.com/karuppaiah/shopping/product/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct1 := &model.Product{
		Name:  "Belts",
		Price: 29.9,
		Stock: 20,
	}
	mockProduct2 := &model.Product{
		Name:  "Belts",
		Price: 29.9,
		Stock: 20,
	}

	mockListProduct := make([]*model.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct1)
	mockListProduct = append(mockListProduct, mockProduct2)
	mockProductRepo.On("Fetch", mock.Anything).Return(mockListProduct, nil)

	u := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

	list, err := u.Fetch(context.TODO())

	assert.NoError(t, err)
	assert.Len(t, list, len(mockListProduct))

	mockProductRepo.AssertExpectations(t)

}

func TestFetchError(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	mockProductRepo.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpexted Error"))

	u := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

	list, err := u.Fetch(context.TODO())

	assert.Error(t, err)
	assert.Len(t, list, 0)
	mockProductRepo.AssertExpectations(t)

}

func TestStore(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := model.Product{
		Name:  "Belts",
		Price: 29.9,
		Stock: 20,
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockProduct := mockProduct
	tempMockProduct.ID = 0

	mockProductRepo.On("Store", mock.Anything, mock.AnythingOfType("*model.Product")).Return(int64(mockProduct.ID), nil)

	u := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockProduct)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockProduct.Name, tempMockProduct.Name)
	mockProductRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockArticle := model.Product{
		Name:  "Belts",
		Price: 29.9,
		Stock: 20,
	}

	mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(true, nil)

	u := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockArticle.ID)

	assert.NoError(t, err)
	assert.True(t, a)
	mockProductRepo.AssertExpectations(t)

}
