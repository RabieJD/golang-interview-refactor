package calculator

import (
	"context"
	"interview/pkg/entity"
)

type CartRepo interface {
	// AddCart adds a New cart to database
	AddCart(ctx context.Context, cart *entity.CartEntity) error
	// UpdateCart updates an existing cart in database, if not existing, it inserts it
	UpdateCart(ctx context.Context, cart *entity.CartEntity) error
	// UpdateItem updates an existing item in database, if not existing, it inserts it
	UpdateItem(ctx context.Context, item *entity.CartItem) error
	// GetCartBySessionID returns a cart by its SessionID
	GetCartBySessionID(ctx context.Context, sessionID string) (*entity.CartEntity, error)
	// GetItemByCartIDAndName returns an item by its CardID and its ProductName
	GetItemByCartIDAndName(ctx context.Context, cartID uint, name string) (*entity.CartItem, error)
	// GetItemsByCartID returns the list of items belonging to a cart
	GetItemsByCartID(ctx context.Context, cartID uint) ([]*entity.CartItem, error)
	// DeleteItem deletes an item by its ID
	DeleteItem(ctx context.Context, itemID uint) error
}

type PriceRepo interface {
	GetPrice(identifier string) (float64, error)
}
