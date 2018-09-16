package repository_test

import (
	"context"
	"testing"

	"github.com/karuppaiah/shopping/model"
	ERepo "github.com/karuppaiah/shopping/promotion/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "sprodid", "sminqty", "dprodid", "dminqty", "disctype", "discount", "priority"}).
		AddRow(1, 2, 2, 4, 0, "D", 15.0, 1).
		AddRow(2, 2, 2, 5, 0, "D", 15.0, 1)

	query := "SELECT id,sprodid,sminqty, dprodid,dminqty,disctype,discount,priority"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := ERepo.NewERepository(db)

	list, err := a.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestStore(t *testing.T) {

	ar := &model.Promotion{
		Sprodid:  2,
		Sminqty:  2,
		Dprodid:  3,
		Dminqty:  0,
		Disctype: "D",
		Discount: 15,
		Priority: 1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rowsSProduct := sqlmock.NewRows([]string{"id", "name", "price", "stock"}).AddRow(2, "Belts", 29.9, 30)
	rowsDProduct := sqlmock.NewRows([]string{"id", "name", "price", "stock"}).AddRow(3, "Belts", 29.9, 30)
	// prodQuery := "SELECT id, name, price, stock FROM products"

	mock.ExpectQuery("SELECT id, name, price, stock FROM products").WillReturnRows(rowsSProduct)
	mock.ExpectQuery("SELECT id, name, price, stock FROM products").WillReturnRows(rowsDProduct)

	rowsPromotion := sqlmock.NewRows([]string{"id", "sprodid", "sminqty", "dprodid", "dminqty", "disctype", "discount", "priority"})

	mock.ExpectQuery("SELECT id, sprodid , sminqty , dprodid , dminqty , disctype , discount,priority").WillReturnRows(rowsPromotion)

	//rows := sqlmock.NewRows([]string{})
	query := "INSERT INTO promotions "
	//query2 := "SELECT id,code,prodid, name,price,items,dprice FROM carts where code = \\? and prodid= \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Sprodid, ar.Sminqty, ar.Dprodid, ar.Dminqty, ar.Disctype, ar.Discount, ar.Priority).WillReturnResult(sqlmock.NewResult(12, 1))

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

	query := "DELETE FROM promotions WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := ERepo.NewERepository(db)

	num := 12
	anDelStatus, err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
	assert.True(t, anDelStatus)
}
