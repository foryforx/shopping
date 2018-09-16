package repository_test

import (
	"context"
	"testing"

	ERepo "github.com/karuppaiah/shopping/cart/repository"
	"github.com/karuppaiah/shopping/model"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "code", "prodid", "name", "price", "items", "dprice"}).
		AddRow(1, "admin", 2, "Belts", 20.0, 2, 0).
		AddRow(1, "admin", 4, "Shoes", 30.0, 5, 0)

	query := "SELECT id,code,prodid, name,price,items,dprice"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := ERepo.NewERepository(db)

	list, err := a.Fetch(context.TODO(), "admin")
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestStore(t *testing.T) {

	ar := &model.Cart{
		Name:   "Belts",
		Price:  29.9,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 0,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rowsProduct := sqlmock.NewRows([]string{"id", "name", "price", "stock"}).AddRow(2, "Belts", 29.9, 30)
	mock.ExpectQuery("SELECT id, name, price, stock FROM products").WillReturnRows(rowsProduct)

	rowsCart := sqlmock.NewRows([]string{"id", "code", "prodid", "name", "price", "items", "dprice"})
	mock.ExpectQuery("SELECT id, code, prodid, name, price, items, dprice").WillReturnRows(rowsCart)

	//rows := sqlmock.NewRows([]string{})
	query := "INSERT  "
	//query2 := "SELECT id,code,prodid, name,price,items,dprice FROM carts where code = \\? and prodid= \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Code, ar.Prodid, ar.Name, ar.Price, ar.Items, ar.Dprice).WillReturnResult(sqlmock.NewResult(12, 1))
	//prep2 := mock.ExpectPrepare(query2)
	//mock.ExpectQuery(query2).WithArgs(ar.Code, ar.Prodid).WillReturnRows(rows)

	a := ERepo.NewERepository(db)

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

	query := "DELETE FROM carts WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := ERepo.NewERepository(db)

	num := 12
	anProductStatus, err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
	assert.True(t, anProductStatus)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// rows := sqlmock.NewRows([]string{"id", "code", "prodid", "name", "price", "items", "dprice"}).
	// 	AddRow(1, "admin", 2, "Belts", 20.0, 2, 0)

	query := "UPDATE carts"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(20.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := ERepo.NewERepository(db)

	ar := &model.Cart{
		ID:     1,
		Name:   "Belts",
		Price:  29.9,
		Items:  20,
		Prodid: 2,
		Code:   "admin",
		Dprice: 20.0,
	}
	aCart, err := a.Update(context.TODO(), ar)
	assert.NoError(t, err)
	assert.NotNil(t, aCart)
	assert.Equal(t, float64(20.0), aCart.Dprice)

}

func TestFetchPromotionDetailsOfCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "sprodid", "sminqty", "dprodid", "dminqty", "disctype", "discount", "priority"}).
		AddRow(1, 4, 2, 2, 0, "P", 15, 1).
		AddRow(1, 2, 2, 2, 0, "F", 15, 1)

	query := "SELECT id,sprodid,sminqty,dprodid,dminqty,disctype,discount,priority"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := ERepo.NewERepository(db)

	list, err := a.FetchPromotionDetailsForCart(context.TODO(), "admin")
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestConvertCartDetailsAsMap(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "code", "prodid", "name", "price", "items", "dprice"}).
		AddRow(1, "admin", 2, "Belts", 20.0, 2, 0).
		AddRow(2, "admin", 4, "Shoes", 30.0, 5, 0)

	query := "SELECT id,code,prodid, name,price,items,dprice"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := ERepo.NewERepository(db)

	mapCart, err := a.ConvertCartDetailsAsMap(context.TODO(), "admin")
	assert.NoError(t, err)
	assert.Len(t, mapCart, 2)
	assert.Equal(t, mapCart[2].Name, "Belts")
	assert.Equal(t, mapCart[4].Name, "Shoes")
}
