package repository

import (
	"fmt"

	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jmoiron/sqlx"
)

type SalesListPostgres struct {
	db *sqlx.DB
}

func NewSalesListPostgres(db *sqlx.DB) *SalesListPostgres {
	return &SalesListPostgres{db: db}
}

func (r *SalesListPostgres) Create(userId int, list rgm.SalesList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, price, amount, total) VALUES ($1, $2, $3, $4) RETURNING id", salesListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Price, list.Amount, (list.Price * float64(list.Amount)))
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SalesListPostgres) GetAll(userId int) ([]rgm.SalesList, error) {
	var lists []rgm.SalesList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.price, tl.amount FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		salesListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *SalesListPostgres) GetAllUserId() ([]int, error) {
	var userIds []int
	query := "SELECT id FROM users"
	err := r.db.Select(&userIds, query)
	return userIds, err
}

func (r *SalesListPostgres) GetUserNameById(id int) (string, error) {
	var userName string
	query := "SELECT name FROM users WHERE id = $1"
	err := r.db.Get(&userName, query, id)
	return userName, err
}

func (r *SalesListPostgres) GetBiggerSale() ([]rgm.SalesRepo, error) {
	var sale []rgm.SalesRepo
	query := "SELECT id, title, amount, price, total FROM sales_lists ORDER BY total DESC LIMIT 3"
	err := r.db.Select(&sale, query)
	return sale, err
}

func (r *SalesListPostgres) GetLowerSale() ([]rgm.SalesRepo, error) {
	var sale []rgm.SalesRepo
	query := "SELECT id, title, amount, price, total FROM sales_lists ORDER BY total ASC LIMIT 3"
	err := r.db.Select(&sale, query)
	return sale, err
}
