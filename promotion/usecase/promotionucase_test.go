package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/karuppaiah/shopping/model"
	"github.com/karuppaiah/shopping/promotion/mocks"
	"github.com/karuppaiah/shopping/promotion/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockPromotionRepo := new(mocks.ERepository)
	mockEnt1 := &model.Promotion{
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "D",
		Discount: 15,
		Priority: 1,
	}
	mockEnt2 := &model.Promotion{
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  5,
		Dminqty:  0,
		Disctype: "D",
		Discount: 15,
		Priority: 1,
	}

	mockListPromotion := make([]*model.Promotion, 0)
	mockListPromotion = append(mockListPromotion, mockEnt1)
	mockListPromotion = append(mockListPromotion, mockEnt2)

	mockPromotionRepo.On("Fetch", mock.Anything).Return(mockListPromotion, nil)

	u := usecase.NewEUsecase(mockPromotionRepo, time.Second*2)

	mockPromotion, err := u.Fetch(context.TODO())

	assert.NoError(t, err)
	assert.NotNil(t, mockPromotion)

	mockPromotionRepo.AssertExpectations(t)

}

func TestFetchError(t *testing.T) {
	mockPromotionRepo := new(mocks.ERepository)

	mockPromotionRepo.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpexted Error"))

	u := usecase.NewEUsecase(mockPromotionRepo, time.Second*2)

	list, err := u.Fetch(context.TODO())

	assert.Error(t, err)
	assert.Nil(t, list)
	mockPromotionRepo.AssertExpectations(t)

}

func TestStore(t *testing.T) {
	mockPromotionRepo := new(mocks.ERepository)
	mockPromotion := model.Promotion{
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "D",
		Discount: 15,
		Priority: 1,
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockPromotion := mockPromotion
	tempMockPromotion.ID = 0

	mockPromotionRepo.On("Store", mock.Anything, mock.AnythingOfType("*model.Promotion")).Return(int64(mockPromotion.ID), nil)

	u := usecase.NewEUsecase(mockPromotionRepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockPromotion)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockPromotion.Sprodid, tempMockPromotion.Sprodid)
	mockPromotionRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockPromotionRepo := new(mocks.ERepository)
	mockPromotion := model.Promotion{
		ID:       1,
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "D",
		Discount: 15,
		Priority: 1,
	}

	mockPromotionRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(true, nil)

	u := usecase.NewEUsecase(mockPromotionRepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockPromotion.ID)

	assert.NoError(t, err)
	assert.True(t, a)
	mockPromotionRepo.AssertExpectations(t)

}
