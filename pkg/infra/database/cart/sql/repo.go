package sql

import (
	"context"
	"gorm.io/gorm"
	"interview/pkg/entity"
	"interview/pkg/infra/database/cart"
	"interview/pkg/infra/database/connection"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (cart.Handler, connection.Migration) {
	r := &repo{db: db}
	return r, r
}

func (r *repo) Migrate() error {
	return r.db.AutoMigrate(&Cart{}, &CartItem{})
}

// AddCart adds a New cart to database
func (r *repo) AddCart(ctx context.Context, cart entity.CartEntity) error {
	dbCart := CartEntityToDBCart(cart)
	result := r.db.WithContext(ctx).Create(&dbCart)
	return result.Error
}

// UpdateCart updates an existing cart in database, if not existing, it inserts it
func (r *repo) UpdateCart(ctx context.Context, cart entity.CartEntity) error {
	dbCart := CartEntityToDBCart(cart)
	result := r.db.WithContext(ctx).Save(&dbCart)
	return result.Error
}

// AddItem adds a new item to database
func (r *repo) AddItem(ctx context.Context, item entity.CartItem) error {
	dbItem := CartItemToDBCartItem(item)
	result := r.db.WithContext(ctx).Create(&dbItem)
	return result.Error
}

// UpdateItem updates an existing item in database, if not existing, it inserts it
func (r *repo) UpdateItem(ctx context.Context, item entity.CartItem) error {
	dbItem := CartItemToDBCartItem(item)
	result := r.db.WithContext(ctx).Save(&dbItem)
	return result.Error
}

// GetCartBySessionID returns a cart by its SessionID
func (r *repo) GetCartBySessionID(ctx context.Context, sessionID string) (*entity.CartEntity, error) {
	var c Cart
	result := r.db.WithContext(ctx).First(&c, "status = ? AND session_id = ?", CartOpen, sessionID)
	if result.Error != nil {
		return nil, result.Error
	}
	entityCart := c.ToCartEntity()
	return &entityCart, nil
}

// GetItemByCartIDAndName returns an item by its CardID and its ProductName
func (r *repo) GetItemByCartIDAndName(ctx context.Context, cartID uint, name string) (*entity.CartItem, error) {
	var i CartItem
	result := r.db.WithContext(ctx).First(&i, "cart_id = ? AND product_name = ?", cartID, name)
	if result.Error != nil {
		return nil, result.Error
	}
	entityItem := i.ToEntityCartItem()
	return &entityItem, nil
}

// GetItemsByCartID returns the list of items belonging to a cart
func (r *repo) GetItemsByCartID(ctx context.Context, cartID uint) ([]*entity.CartItem, error) {
	var items []*CartItem
	result := r.db.WithContext(ctx).Find(&items, "cart_id = ?", cartID)
	if result.Error != nil {
		return nil, result.Error
	}
	var entityItems []*entity.CartItem
	for _, i := range items {
		entityItem := i.ToEntityCartItem()
		entityItems = append(entityItems, &entityItem)
	}
	return entityItems, nil
}

// DeleteItem deletes an item by its ID
func (r *repo) DeleteItem(ctx context.Context, itemID uint) error {
	result := r.db.WithContext(ctx).Delete(&CartItem{}, "ID = ?", itemID)
	return result.Error
}
