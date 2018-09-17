package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/karuppaiah/shopping/login/mocks"
	"github.com/karuppaiah/shopping/login/usecase"
	"github.com/karuppaiah/shopping/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchUser(t *testing.T) {
	mockLoginRepo := new(mocks.ERepository)
	mockEnt1 := &model.Login{
		Username: "kal",
		Password: "kal",
	}
	mockListLogin := make([]*model.Login, 0)
	mockListLogin = append(mockListLogin, mockEnt1)

	mockLoginRepo.On("FetchLoginwithUsername", mock.Anything, mock.AnythingOfType("string")).Return(mockListLogin, nil)

	u := usecase.NewEUsecase(mockLoginRepo, time.Second*2)

	mockLogin, err := u.Fetch(context.TODO(), "kal")

	assert.NoError(t, err)
	assert.NotNil(t, mockLogin)
	assert.Equal(t, "kal", mockLogin[0].Username)
	mockLoginRepo.AssertExpectations(t)

}

func TestFetchError(t *testing.T) {
	mockLoginRepo := new(mocks.ERepository)

	mockLoginRepo.On("FetchLoginwithUsername", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpexted Error"))

	u := usecase.NewEUsecase(mockLoginRepo, time.Second*2)

	list, err := u.Fetch(context.TODO(), "kal")

	assert.Error(t, err)
	assert.Nil(t, list)
	mockLoginRepo.AssertExpectations(t)

}

func TestStore(t *testing.T) {
	mockLoginRepo := new(mocks.ERepository)
	mockLogin := model.Login{
		Username: "kal",
		Password: "kal",
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockLogin := mockLogin

	mockLoginRepo.On("Store", mock.Anything, mock.AnythingOfType("*model.Login")).Return(int64(1), nil)

	u := usecase.NewEUsecase(mockLoginRepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockLogin)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockLogin.Username, tempMockLogin.Username)
	assert.Equal(t, int64(1), a)
	mockLoginRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockLoginRepo := new(mocks.ERepository)
	mockLogin := model.Login{
		Username: "kal",
		Password: "kal",
	}

	mockLoginRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(true, nil)

	u := usecase.NewEUsecase(mockLoginRepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockLogin.Username)

	assert.NoError(t, err)
	assert.True(t, a)
	mockLoginRepo.AssertExpectations(t)

}
