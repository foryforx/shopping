package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/karuppaiah/shopping/cart"
	"github.com/karuppaiah/shopping/model"
)

type eRepository struct {
	Conn *sql.DB
}

func NewERepository(conn *sql.DB) cart.ERepository {

	return &eRepository{conn}
}

func (m *eRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Cart, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Cart, 0)
	for rows.Next() {
		t := new(model.Cart)

		err = rows.Scan(
			&t.ID,
			&t.Code,
			&t.Prodid,
			&t.Name,
			&t.Price,
			&t.Items,
			&t.Dprice,
		)

		if err != nil {

			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *eRepository) fetchProduct(ctx context.Context, query string, args ...interface{}) ([]*model.Product, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Product, 0)
	for rows.Next() {
		t := new(model.Product)

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Price,
			&t.Stock,
		)

		if err != nil {

			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *eRepository) Fetch(ctx context.Context, user string) ([]*model.Cart, error) {

	query := `SELECT id,code,prodid, name,price,items,dprice
  						FROM carts where code = ?`
	return m.fetch(ctx, query, user)

}

func (m *eRepository) Store(ctx context.Context, a *model.Cart) (int64, error) {
	// verifyQuery := `SELECT id,code,prodid, name,price,items,dprice
	// FROM carts where code = ? and prodid= ?`
	// list, verifyerr := m.fetch(ctx, verifyQuery, a.Code, a.Prodid)
	// fmt.Println(verifyerr)
	// if verifyerr != nil {
	// 	return 0, errors.New("Invalid request")
	// }
	// if len(list) > 0 {
	// 	return 0, errors.New("Item already exists in cart")
	// }
	// get the prooname and price from product
	prodQuery := `SELECT id, name,price, stock
						  FROM products where id = ?`
	prodItems, err := m.fetchProduct(ctx, prodQuery, a.Prodid)

	if err != nil {
		return 0, err
	}

	if len(prodItems) <= 0 {
		lenErr := fmt.Errorf("Item doesnt exists")
		return 0, lenErr
	}
	a.Name = prodItems[0].Name
	a.Price = prodItems[0].Price
	fmt.Println("model:", a.Code, a.Prodid)
	cartQuery := `SELECT id, code, prodid, name, price, items, dprice
	FROM carts where code = ? and prodid = ?`
	cartItems, cerr := m.fetch(ctx, cartQuery, a.Code, a.Prodid)
	fmt.Println("modely:", len(cartItems), a.Code, a.Prodid, cerr)
	if cerr != nil {
		return 0, cerr
	}

	if len(cartItems) > 0 {
		lencErr := fmt.Errorf("Item already exists in cart. Please delete and add again")
		return 0, lencErr
	}

	query := `INSERT INTO carts ( code , prodid , name , price , items , dprice ) VALUES ( ? , ? , ? , ? , ? , ? )`

	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, a.Code, a.Prodid, a.Name, a.Price, a.Items, a.Dprice)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *eRepository) Delete(ctx context.Context, id int) (bool, error) {
	query := "DELETE FROM carts WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	res, err := stmt.ExecContext(ctx, id)
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
