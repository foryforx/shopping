package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/karuppaiah/shopping/model"
	product "github.com/karuppaiah/shopping/product"
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) product.ProductRepository {

	return &productRepository{conn}
}

func (m *productRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Product, error) {
	fmt.Println("In Repo")
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	fmt.Println(query, rows)
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

func (m *productRepository) Fetch(ctx context.Context) ([]*model.Product, error) {

	query := `SELECT id,name,price,stock
  						FROM products`

	return m.fetch(ctx, query)

}

func (m *productRepository) Store(ctx context.Context, a *model.Product) (int64, error) {

	query := `INSERT INTO products (name,price,stock) VALUES(? , ? , ?)`

	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, a.Name, a.Price, a.Stock)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (m *productRepository) Delete(ctx context.Context, id int) (bool, error) {
	query := "DELETE FROM products WHERE id = ?"
	fmt.Println("id:", id)

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
