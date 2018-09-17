package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/karuppaiah/shopping/model"
	productRepo "github.com/karuppaiah/shopping/product/repository"
	"github.com/karuppaiah/shopping/promotion"
)

type eRepository struct {
	Conn *sql.DB
}

func NewERepository(conn *sql.DB) promotion.ERepository {

	return &eRepository{conn}
}

func (m *eRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Promotion, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Promotion, 0)
	for rows.Next() {
		t := new(model.Promotion)

		err = rows.Scan(
			&t.ID,
			&t.Sprodid,
			&t.Sminqty,
			&t.Dprodid,
			&t.Dminqty,
			&t.Disctype,
			&t.Discount,
			&t.Priority,
		)
		fmt.Println("prodid:", t.Sprodid)
		if err != nil {

			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

// func (m *eRepository) fetchProduct(ctx context.Context, query string, args ...interface{}) ([]*model.Product, error) {

// 	rows, err := m.Conn.QueryContext(ctx, query, args...)

// 	if err != nil {

// 		return nil, err
// 	}
// 	defer rows.Close()
// 	result := make([]*model.Product, 0)
// 	for rows.Next() {
// 		t := new(model.Product)

// 		err = rows.Scan(
// 			&t.ID,
// 			&t.Name,
// 			&t.Price,
// 			&t.Stock,
// 		)

// 		if err != nil {

// 			return nil, err
// 		}

// 		result = append(result, t)
// 	}

// 	return result, nil
// }
// Fetch promotions
func (m *eRepository) Fetch(ctx context.Context) ([]*model.Promotion, error) {

	query := `SELECT id,sprodid,sminqty, dprodid,dminqty,disctype,discount,priority
  						FROM promotions`
	return m.fetch(ctx, query)

}

// Fetch promotion with query for other repository
func (m *eRepository) FetchPromotionwithQuery(ctx context.Context, query string) ([]*model.Promotion, error) {
	fmt.Println("Promotion Repo")
	return m.fetch(ctx, query)

}

func (m *eRepository) Store(ctx context.Context, a *model.Promotion) (int64, error) {
	prR := productRepo.NewProductRepository(m.Conn)
	// Check source product
	prodQuery := "SELECT id, name, price, stock FROM products where id = '" + strconv.Itoa(a.Sprodid) + "'"
	sprodItems, serr := prR.FetchProductWithQuery(ctx, prodQuery)
	// sprodItems, serr := m.fetchProduct(ctx, prodQuery, a.Sprodid)

	if serr != nil {
		return 0, serr
	}

	if len(sprodItems) <= 0 {
		slenErr := fmt.Errorf("Source Item doesnt exists")
		return 0, slenErr
	}
	//Check destincation product
	//dprodQuery := `SELECT id, name,price,stock FROM products where id = ?`
	// dprodItems, derr := m.fetchProduct(ctx, prodQuery, a.Dprodid)
	prodQuery = "SELECT id, name, price, stock FROM products where id = '" + strconv.Itoa(a.Dprodid) + "'"

	dprodItems, derr := prR.FetchProductWithQuery(ctx, prodQuery)
	if derr != nil {
		return 0, derr
	}

	if len(dprodItems) <= 0 {
		dlenErr := fmt.Errorf("DEstination Item doesnt exists")
		return 0, dlenErr
	}
	// if promotion items is already available, return error
	promotionExistQuery := `SELECT id, sprodid , sminqty , dprodid , dminqty , disctype , discount,priority
	FROM promotions where sprodid = ? and dprodid = ?`
	promtionItems, perr := m.fetch(ctx, promotionExistQuery, a.Sprodid, a.Dprodid)
	fmt.Println("modely:", len(promtionItems), a.Sprodid, a.Dprodid, perr)
	if perr != nil {
		return 0, perr
	}

	if len(promtionItems) > 0 {
		//lencErr := fmt.Errorf("Item already exists in promotion. Please delete and add again")
		fmt.Println("Item already exists in promotion. Please delete and add again ")
		return 0, errors.New("Item already exists in promotion. Please delete and add again ")
	}
	// add promotion item
	query := `INSERT INTO promotions ( sprodid , sminqty , dprodid , dminqty , disctype , discount,priority ) VALUES ( ? , ? , ? , ? , ? , ? , ?)`

	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, a.Sprodid, a.Sminqty, a.Dprodid, a.Dminqty, a.Disctype, a.Discount, a.Priority)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Delete of promotion item
func (m *eRepository) Delete(ctx context.Context, id int) (bool, error) {
	query := "DELETE FROM promotions WHERE id = ?"

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
