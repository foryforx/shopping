package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/karuppaiah/shopping/cart"
	"github.com/karuppaiah/shopping/model"
)

//eRepository with Connection to DB
type eRepository struct {
	Conn *sql.DB
}

// To create new Repository
func NewERepository(conn *sql.DB) cart.ERepository {

	return &eRepository{conn}
}

// Fetch cart item for arguments and return as list
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

// Fetch product for validations and return as list of product
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

// Fetch promotion for validation and calculation and return as list of promotion
func (m *eRepository) fetchPromotion(ctx context.Context, query string, args ...interface{}) ([]*model.Promotion, error) {

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

		if err != nil {

			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

// Fetch cart items in list format for the user
func (m *eRepository) Fetch(ctx context.Context, user string) ([]*model.Cart, error) {

	query := `SELECT id,code,prodid, name,price,items,dprice
  						FROM carts where code = ?`
	return m.fetch(ctx, query, user)

}

//Get cart details as map
func (m *eRepository) ConvertCartDetailsAsMap(ctx context.Context, user string) (map[int](*model.Cart), error) {
	query := `SELECT id,code,prodid, name,price,items,dprice FROM carts where code = ?`
	cartList, err := m.fetch(ctx, query, user)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	fmt.Println("Cartlist:", cartList)
	var mapCart = map[int]*model.Cart{}
	for i := 0; i < len(cartList); i++ {
		mapCart[cartList[i].Prodid] = cartList[i]
	}
	return mapCart, nil
}

// Get promotion detauls as list
func (m *eRepository) FetchPromotionDetailsForCart(ctx context.Context, user string) ([]*model.Promotion, error) {
	query := `SELECT id,sprodid,sminqty,dprodid,dminqty,disctype,discount,priority FROM promotions where sprodid in (select prodid from carts where code = ?) order by priority desc`
	return m.fetchPromotion(ctx, query, user)
}

// //
// func (m *eRepository) RefreshCart(ctx context.Context, user string) error {
// 	mCart, err := m.ConvertCartDetailsAsMap(ctx, user)
// 	fmt.Println("Refresh:", mCart)
// 	if err != nil {
// 		return err
// 	}
// 	lPromo, err := m.FetchPromotionDetailsForCart(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	for i := 0; i < len(lPromo); i++ {
// 		fmt.Println(i, ".Source product id present:", mCart[lPromo[i].Sprodid])
// 		sCartItem := mCart[lPromo[i].Sprodid]
// 		dCartItem := mCart[lPromo[i].Dprodid]
// 		fmt.Println("Source cart item:", sCartItem)
// 		fmt.Println("Destination Cart items:", dCartItem)

// 		if sCartItem.Items >= lPromo[i].Sminqty && dCartItem != nil {
// 			fmt.Println("Promotion applicable")
// 			noOfItemDiscApplied := 0
// 			if dCartItem.Items > lPromo[i].Dminqty {
// 				noOfItemDiscApplied = dCartItem.Items - lPromo[i].Dminqty
// 				if lPromo[i].Disctype == "P" {
// 					dCartItem.Dprice = ((lPromo[i].Discount / 100) * dCartItem.Price) * float64(noOfItemDiscApplied)
// 				} else if lPromo[i].Disctype == "F" {
// 					dCartItem.Dprice = (lPromo[i].Discount) * float64(noOfItemDiscApplied)
// 				}
// 				fmt.Println("Final cart item:", dCartItem)
// 				m.Update(ctx, dCartItem)
// 			}

// 		}

// 	}
// 	return nil
// }
//END move as BL

// Update the cart items discount
func (m *eRepository) Update(ctx context.Context, ar *model.Cart) (*model.Cart, error) {
	query := `UPDATE carts set dprice = ? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, ar.Dprice, ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}

	return ar, nil
}

// Add a new cart item
func (m *eRepository) Store(ctx context.Context, a *model.Cart) (int64, error) {

	prodQuery := `SELECT id, name, price, stock FROM products where id = ?`
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
	if a.Items > prodItems[0].Stock {
		stockErr := fmt.Errorf("Stock too high")
		fmt.Println("Stock too much requested")
		return 0, stockErr
	}
	fmt.Println("model:", a.Code, a.Prodid)
	cartQuery := `SELECT id, code, prodid, name, price, items, dprice FROM carts where code = ? and prodid = ?`
	cartItems, cerr := m.fetch(ctx, cartQuery, a.Code, a.Prodid)
	fmt.Println("modely:", len(cartItems), a.Code, a.Prodid, cerr)
	if cerr != nil {
		return 0, cerr
	}

	if len(cartItems) > 0 {
		//lencErr := fmt.Errorf("Item already exists in cart. Please delete and add again")
		fmt.Println("Item already exists in cart. Please delete and add again ")
		return 0, errors.New("Item already exists in cart. Please delete and add again ")
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

// Delete a cart item
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
