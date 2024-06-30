package sql

import (
	"gorm.io/gorm"
	"interview/pkg/entity"
	"interview/pkg/infra/database/cart/sql"
	"time"
)

// TTToCartEntity is a test case for a ToCartEntity  function
type TTToCartEntity struct {
	Input    sql.Cart
	Expected entity.CartEntity
}

// GenerateTTToCartEntity  generates test cases for the ToCartEntity  function
func GenerateTTToCartEntity() []TTToCartEntity {
	t := time.Now()
	return []TTToCartEntity{
		{
			Input: sql.Cart{
				Model:     gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				Total:     100.0,
				SessionID: "session1",
				Status:    sql.CartOpen,
			},
			Expected: entity.CartEntity{
				Model:     gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				Total:     100.0,
				SessionID: "session1",
				Status:    "open",
			},
		},
		{
			Input: sql.Cart{
				Model:     gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				Total:     250.0,
				SessionID: "session2",
				Status:    sql.CartClosed,
			},
			Expected: entity.CartEntity{
				Model:     gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				Total:     250.0,
				SessionID: "session2",
				Status:    "closed",
			},
		},
	}
}

// TTCartEntityToDBCart  is a test case for a CartEntityToDBCart function
type TTCartEntityToDBCart struct {
	Input    entity.CartEntity
	Expected sql.Cart
}

// GenerateTTCartEntityToDBCart generates test cases for the CartEntityToDBCart function
func GenerateTTCartEntityToDBCart() []TTCartEntityToDBCart {
	t := time.Now()
	return []TTCartEntityToDBCart{
		{
			Input: entity.CartEntity{
				Model:     gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				Total:     200.0,
				SessionID: "session2",
				Status:    "closed",
			},
			Expected: sql.Cart{
				Model:     gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				Total:     200.0,
				SessionID: "session2",
				Status:    sql.CartClosed,
			},
		},
		{
			Input: entity.CartEntity{
				Model:     gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				Total:     150.0,
				SessionID: "session3",
				Status:    "open",
			},
			Expected: sql.Cart{
				Model:     gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				Total:     150.0,
				SessionID: "session3",
				Status:    sql.CartOpen,
			},
		},
	}
}

// TTToEntityCartItem is a test case for a ToEntityCartItem  function
type TTToEntityCartItem struct {
	Input    sql.CartItem
	Expected entity.CartItem
}

// GenerateTTToEntityCartItem generates test cases for the ToEntityCartItem  function
func GenerateTTToEntityCartItem() []TTToEntityCartItem {
	t := time.Now()
	return []TTToEntityCartItem{
		{
			Input: sql.CartItem{
				Model:       gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
			Expected: entity.CartItem{
				Model:       gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
		},
		{
			Input: sql.CartItem{
				Model:       gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				ProductName: "Product2",
				Quantity:    1,
				Price:       20.0,
			},
			Expected: entity.CartItem{
				Model:       gorm.Model{ID: 2, CreatedAt: t, UpdatedAt: t},
				ProductName: "Product2",
				Quantity:    1,
				Price:       20.0,
			},
		},
	}

}

// TTCartItemToDBCartItem is a test case for a CartItemToDBCartItem  function
type TTCartItemToDBCartItem struct {
	Input    entity.CartItem
	Expected sql.CartItem
}

// GenerateTTCartItemToDBCartItem generates test cases for the CartItemToDBCartItem  function
func GenerateTTCartItemToDBCartItem() []TTCartItemToDBCartItem {
	t := time.Now()
	return []TTCartItemToDBCartItem{
		{
			Input: entity.CartItem{
				Model:       gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				CartID:      1,
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
			Expected: sql.CartItem{

				Model:       gorm.Model{ID: 1, CreatedAt: t, UpdatedAt: t},
				CartID:      1,
				ProductName: "Product1",
				Quantity:    2,
				Price:       10.0,
			},
		},
	}
}
