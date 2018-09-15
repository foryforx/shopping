package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/karuppaiah/shopping/cart/mocks"
	"github.com/karuppaiah/shopping/cart/usecase"
	"github.com/karuppaiah/shopping/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockEnt1 := &model.Cart{
		Name:   "New Belt",
		Price:  29.9,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt2 := &model.Cart{
		Name:   "New Belt",
		Price:  29.9,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}

	mockListCart := make([]*model.Cart, 0)
	mockListCart = append(mockListCart, mockEnt1)
	mockListCart = append(mockListCart, mockEnt2)

	mockCartRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListCart, nil)

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	mockCart, err := u.Fetch(context.TODO(), "admin")

	assert.NoError(t, err)
	assert.NotNil(t, mockCart)

	mockCartRepo.AssertExpectations(t)

}

func TestFetchError(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)

	mockCartRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpexted Error"))

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	list, err := u.Fetch(context.TODO(), "admin")

	assert.Error(t, err)
	assert.Nil(t, list)
	mockCartRepo.AssertExpectations(t)

}

func TestStore(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockCart := model.Cart{
		Name:   "New Belt",
		Price:  29.9,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockCart := mockCart
	tempMockCart.ID = 0

	mockCartRepo.On("Store", mock.Anything, mock.AnythingOfType("*model.Cart")).Return(int64(mockCart.ID), nil)

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockCart)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockCart.Name, tempMockCart.Name)
	mockCartRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockCart := model.Cart{
		Name:   "New Belt",
		Price:  29.9,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}

	mockCartRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(true, nil)

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockCart.ID)

	assert.NoError(t, err)
	assert.True(t, a)
	mockCartRepo.AssertExpectations(t)

}
