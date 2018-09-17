package repository_test

import (
	"context"
	"testing"

	ERepo "github.com/karuppaiah/shopping/login/repository"
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
	rows := sqlmock.NewRows([]string{"username", "password"}).
		AddRow("kal", "kal")

	query := "SELECT username,password FROM"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := ERepo.NewERepository(db)

	list, err := a.FetchLoginwithUsername(context.TODO(), "kal")
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestStore(t *testing.T) {

	ar := &model.Login{
		Username: "kal",
		Password: "kal",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rowsLogin := sqlmock.NewRows([]string{"username", "password"})

	mock.ExpectQuery("SELECT username,password FROM").WillReturnRows(rowsLogin)

	//rows := sqlmock.NewRows([]string{})
	query := "INSERT INTO logins "
	//query2 := "SELECT id,code,prodid, name,price,items,dprice FROM carts where code = \\? and prodid= \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Username, ar.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	//mock.ExpectQuery(query2).WithArgs(ar.Code, ar.Prodid).WillReturnRows(rows)

	a := ERepo.NewERepository(db)

	lastId, err := a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastId)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM logins WHERE"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("kal").WillReturnResult(sqlmock.NewResult(1, 1))

	a := ERepo.NewERepository(db)

	anDelStatus, err := a.Delete(context.TODO(), "kal")
	assert.NoError(t, err)
	assert.True(t, anDelStatus)
}
