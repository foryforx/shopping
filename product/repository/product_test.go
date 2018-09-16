package repository_test

import (
	"context"
	"testing"

	"github.com/karuppaiah/shopping/model"
	productRepo "github.com/karuppaiah/shopping/product/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
		AddRow(1, "Belts", 20.0, 10).
		AddRow(2, "Shirts", 60, 5)

	query := "SELECT id,name,price,stock"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := productRepo.NewProductRepository(db)

	list, err := a.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestStore(t *testing.T) {

	ar := &model.Product{
		Name:  "New Belt",
		Price: 29.9,
		Stock: 20,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO products "
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Price, ar.Stock).WillReturnResult(sqlmock.NewResult(12, 1))

	a := productRepo.NewProductRepository(db)

	lastId, err := a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), lastId)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM products WHERE "

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := productRepo.NewProductRepository(db)

	num := 12
	anProductStatus, err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
	assert.True(t, anProductStatus)
}
