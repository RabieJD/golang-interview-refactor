package controllers

import "context"

// data keys
const (
	DataErrorKey     = "error"
	DataCartItemsKey = "CartItems"
)

// item keys
const (
	CartDataIDKey       = "ID"
	CartDataQuantityKey = "Quantity"
	CartDataPriceKey    = "Price"
	CartDataProductKey  = "Product"
)

type CalculatorHandler interface {
	AddItem(ctx context.Context, item CartItemForm, sessionID string) (code int, err error)
	GetCartData(ctx context.Context, sessionID string) (items []CartData, err error)
	DeleteItemFromCart(ctx context.Context, itemID string, sessionID string) (code int, err error)
}

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}

type CartData map[string]interface{}
