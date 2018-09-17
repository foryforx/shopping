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
		ID:     1,
		Name:   "New Belt",
		Price:  30,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt2 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt3 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 300,
	}
	mockEntMap := map[int](*model.Cart){}
	mockEntMap[2] = mockEnt1
	mockEntMap[3] = mockEnt2
	mockListCart := make([]*model.Cart, 0)
	mockListCart = append(mockListCart, mockEnt1)
	mockListCart = append(mockListCart, mockEnt2)

	mockPromo1 := &model.Promotion{
		ID:       1,
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "P",
		Discount: 15,
		Priority: 1,
	}

	mockListPromotion := make([]*model.Promotion, 0)
	mockListPromotion = append(mockListPromotion, mockPromo1)

	mockCartRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListCart, nil)
	mockCartRepo.On("ConvertCartDetailsAsMap", mock.Anything, mock.Anything).Return(mockEntMap, nil)
	mockCartRepo.On("FetchPromotionDetailsForCart", mock.Anything, mock.Anything).Return(mockListPromotion, nil)
	mockCartRepo.On("Update", mock.Anything, mock.Anything).Return(mockEnt3, nil)

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	mockCart, err := u.Fetch(context.TODO(), "admin")

	assert.NoError(t, err)
	assert.NotNil(t, mockCart)
	assert.Equal(t, 300.0, mockEntMap[3].Dprice)
	mockCartRepo.AssertExpectations(t)

}

func TestFetchError(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockEnt1 := &model.Cart{
		ID:     1,
		Name:   "New Belt",
		Price:  29.9,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt2 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt3 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 300,
	}
	mockEntMap := map[int](*model.Cart){}
	mockEntMap[2] = mockEnt1
	mockEntMap[3] = mockEnt2
	mockListCart := make([]*model.Cart, 0)
	mockListCart = append(mockListCart, mockEnt1)
	mockListCart = append(mockListCart, mockEnt2)

	mockPromo1 := &model.Promotion{
		ID:       1,
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "P",
		Discount: 15,
		Priority: 1,
	}

	mockListPromotion := make([]*model.Promotion, 0)
	mockListPromotion = append(mockListPromotion, mockPromo1)
	mockCartRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpexted Error"))
	mockCartRepo.On("ConvertCartDetailsAsMap", mock.Anything, mock.Anything).Return(mockEntMap, nil)
	mockCartRepo.On("FetchPromotionDetailsForCart", mock.Anything, mock.Anything).Return(mockListPromotion, nil)
	mockCartRepo.On("Update", mock.Anything, mock.Anything).Return(mockEnt3, nil)

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

func TestGetCartTotalValue(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockEnt1 := &model.Cart{
		ID:     1,
		Name:   "New Belt",
		Price:  30,
		Items:  2,
		Prodid: 2,
		Code:   "admin",
		Dprice: 30,
	}
	mockEnt2 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}
	mockListCart := make([]*model.Cart, 0)
	mockListCart = append(mockListCart, mockEnt1)
	mockListCart = append(mockListCart, mockEnt2)
	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	a := u.GetTotalCartValue(mockListCart)

	assert.Equal(t, 2030.0, a)
	mockCartRepo.AssertExpectations(t)

}

func TestFetchReevaluateDiscount(t *testing.T) {
	mockCartRepo := new(mocks.ERepository)
	mockEnt1 := &model.Cart{
		ID:     1,
		Name:   "New Belt",
		Price:  30,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 400,
	}
	mockEnt2 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 0,
	}
	mockEnt3 := &model.Cart{
		ID:     2,
		Name:   "New Belt",
		Price:  100,
		Items:  20,
		Prodid: 3,
		Code:   "admin",
		Dprice: 300,
	}
	mockEnt4 := &model.Cart{
		ID:     1,
		Name:   "New Belt",
		Price:  30,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 0,
	}
	mockEntMap := map[int](*model.Cart){}
	mockEntMap[2] = mockEnt1
	mockEntMap[3] = mockEnt2
	mockListCart := make([]*model.Cart, 0)
	mockListCart = append(mockListCart, mockEnt1)
	mockListCart = append(mockListCart, mockEnt2)

	mockPromo1 := &model.Promotion{
		ID:       1,
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "P",
		Discount: 15,
		Priority: 1,
	}

	mockListPromotion := make([]*model.Promotion, 0)
	mockListPromotion = append(mockListPromotion, mockPromo1)

	mockCartRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListCart, nil)
	mockCartRepo.On("ConvertCartDetailsAsMap", mock.Anything, mock.Anything).Return(mockEntMap, nil)
	mockCartRepo.On("FetchPromotionDetailsForCart", mock.Anything, mock.Anything).Return(mockListPromotion, nil)
	mockCartRepo.On("Update", mock.Anything, mockEnt2).Return(mockEnt3, nil)
	mockCartRepo.On("Update", mock.Anything, mockEnt1).Return(mockEnt4, nil)

	u := usecase.NewEUsecase(mockCartRepo, time.Second*2)

	mockCart, err := u.Fetch(context.TODO(), "admin")

	assert.NoError(t, err)
	assert.NotNil(t, mockCart)
	assert.Equal(t, 300.0, mockEntMap[3].Dprice)
	assert.Equal(t, float64(0), mockEntMap[2].Dprice)
	mockCartRepo.AssertExpectations(t)

}
