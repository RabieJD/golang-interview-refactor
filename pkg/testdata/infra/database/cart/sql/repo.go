package sql

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"interview/pkg/entity"
	"interview/pkg/infra/database/cart/sql"
)

type TestCase struct {
	Name     string
	Mock     func(mock sqlmock.Sqlmock)
	Args     interface{}
	HasError bool
}

func GenerateADDCartTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `carts`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 100.0, "session1", "open").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Args: &entity.CartEntity{
				Total:     100.0,
				SessionID: "session1",
				Status:    "open",
			},
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `carts`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 100.0, "session1", "open").
					WillReturnError(errors.New("insert error"))
				mock.ExpectRollback()
			},
			Args: &entity.CartEntity{
				Total:     100.0,
				SessionID: "session1",
				Status:    "open",
			},
			HasError: true,
		},
	}
}

func GenerateUpdateCartTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE `carts`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 150.0, "session1", "closed", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Args: &entity.CartEntity{
				Model:     gorm.Model{ID: 1},
				Total:     150.0,
				SessionID: "session1",
				Status:    "closed",
			},
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE `carts`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 150.0, "session1", "closed", 1).
					WillReturnError(errors.New("update error"))
				mock.ExpectRollback()
			},
			Args: &entity.CartEntity{
				Model:     gorm.Model{ID: 1},
				Total:     150.0,
				SessionID: "session1",
				Status:    "closed",
			},

			HasError: true,
		},
	}
}

func GenerateAddItemTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `cart_items`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 0, "Product1", 2, 10.0).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Args: &entity.CartItem{
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `cart_items`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 0, "Product1", 2, 10.0).
					WillReturnError(errors.New("insert error"))
				mock.ExpectRollback()
			},
			Args: &entity.CartItem{
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
			HasError: true,
		},
	}
}

func GenerateUpdateItemTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE `cart_items`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 0, "Product1", 3, 12.0, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Args: &entity.CartItem{
				Model:       gorm.Model{ID: 1},
				ProductName: "Product1",
				Quantity:    3,
				Price:       12.0,
			},
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE `cart_items`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 0, "Product1", 3, 12.0, 1).
					WillReturnError(errors.New("update error"))
				mock.ExpectRollback()
			},
			Args: &entity.CartItem{
				Model:       gorm.Model{ID: 1},
				ProductName: "Product1",
				Quantity:    3,
				Price:       12.0,
			},
			HasError: true,
		},
	}
}

func GenerateGetCartBySessionIDTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "total", "session_id", "status"}).
					AddRow(1, 100.0, "session1", "open")
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(sql.CartOpen, "session1").
					WillReturnRows(rows)
			},
			Args:     "session1",
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(sql.CartOpen, "session1").
					WillReturnError(errors.New("query error"))
			},
			Args:     "session1",
			HasError: true,
		},
	}
}

func GenerateGetItemByCartIDAndNameTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "product_name", "quantity", "price"}).
					AddRow(1, "Product1", 2, 10.0)
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(uint(1), "Product1").
					WillReturnRows(rows)
			},
			Args:     []interface{}{uint(1), "Product1"},
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(uint(1), "Product1").
					WillReturnError(errors.New("query error"))
			},
			Args:     []interface{}{uint(1), "Product1"},
			HasError: true,
		},
	}
}

func GenerateGetItemsByCartIDTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "product_name", "quantity", "price"}).
					AddRow(1, "Product1", 2, 10.0).
					AddRow(2, "Product2", 1, 20.0)
				mock.ExpectQuery("SELECT \\* FROM `cart_items` WHERE cart_id = \\?").
					WithArgs(uint(1)).
					WillReturnRows(rows)
			},
			Args:     uint(1),
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT \\* FROM `cart_items` WHERE cart_id = \\?").
					WithArgs(uint(1)).
					WillReturnError(errors.New("query error"))
			},
			Args:     uint(1),
			HasError: true,
		},
	}
}

func GenerateDeleteItemTT() []TestCase {
	return []TestCase{
		{
			Name: "Success",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE (.+)").
					WithArgs(sqlmock.AnyArg(), uint(1)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Args:     uint(1),
			HasError: false,
		},
		{
			Name: "Fail",
			Mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE (.+)").
					WithArgs(sqlmock.AnyArg(), uint(1)).
					WillReturnError(errors.New("delete error"))
				mock.ExpectRollback()
			},
			Args:     uint(1),
			HasError: true,
		},
	}
}
