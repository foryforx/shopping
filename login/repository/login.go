package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/karuppaiah/shopping/login"
	"github.com/karuppaiah/shopping/model"
)

type eRepository struct {
	Conn *sql.DB
}

func NewERepository(conn *sql.DB) login.ERepository {

	return &eRepository{conn}
}

// Get the logins based on query and params
func (m *eRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Login, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Login, 0)
	for rows.Next() {
		t := new(model.Login)

		err = rows.Scan(
			&t.Username,
			&t.Password,
		)

		if err != nil {

			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

// Fetch all logins
func (m *eRepository) Fetch(ctx context.Context) ([]*model.Login, error) {

	query := `SELECT username,password FROM logins`
	return m.fetch(ctx, query)

}

// Fetch login with query for other repository
func (m *eRepository) FetchLoginwithUsername(ctx context.Context, user string) ([]*model.Login, error) {
	query := `SELECT username,password FROM logins where username = ?`

	return m.fetch(ctx, query, user)

}

// Add login
func (m *eRepository) Store(ctx context.Context, a *model.Login) (int64, error) {

	// if login items is already available, return error
	loginExistQuery := `SELECT username,password FROM logins where username = ?`
	loginItems, perr := m.fetch(ctx, loginExistQuery, a.Username)
	if perr != nil {
		return 0, perr
	}

	if len(loginItems) > 0 {
		//lencErr := fmt.Errorf("Item already exists in login. Please delete and add again")
		fmt.Println("Item already exists in login. Choose another username")
		return 0, errors.New("Item already exists. Please choose another username ")
	}
	// add login item
	query := `INSERT INTO logins ( username, password ) VALUES ( ? , ? )`

	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, a.Username, a.Password)

	if err != nil {
		return 0, err
	}
	// Return nmmber of rows inserted
	return res.RowsAffected()
}

// Delete of login item
func (m *eRepository) Delete(ctx context.Context, username string) (bool, error) {
	query := "DELETE FROM logins WHERE username = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	res, err := stmt.ExecContext(ctx, username)
	if err != nil {

		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)

		return false, err
	}

	return true, nil
}
