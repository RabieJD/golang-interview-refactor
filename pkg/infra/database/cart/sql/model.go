package sql

import "C"
import (
	"gorm.io/gorm"
	"interview/pkg/entity"
)

type CartStatus string

const (
	CartOpen   CartStatus = "open"
	CartClosed CartStatus = "closed"
)

type Cart struct {
	gorm.Model
	Total     float64
	SessionID string
	Status    CartStatus
}

func (c Cart) ToCartEntity() entity.CartEntity {
	return entity.CartEntity{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: c.DeletedAt,
		},
		Total:     c.Total,
		SessionID: c.SessionID,
		Status:    string(c.Status),
	}
}

func CartEntityToDBCart(c entity.CartEntity) Cart {
	return Cart{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: c.DeletedAt,
		},
		Total:     c.Total,
		SessionID: c.SessionID,
		Status:    CartStatus(c.Status),
	}
}

type CartItem struct {
	gorm.Model
	CartID      uint
	ProductName string
	Quantity    int
	Price       float64
}

func (ci CartItem) ToEntityCartItem() entity.CartItem {
	return entity.CartItem{
		Model: gorm.Model{
			ID:        ci.ID,
			CreatedAt: ci.CreatedAt,
			UpdatedAt: ci.UpdatedAt,
			DeletedAt: ci.DeletedAt,
		},
		CartID:      ci.CartID,
		ProductName: ci.ProductName,
		Quantity:    ci.Quantity,
		Price:       ci.Price,
	}
}

func CartItemToDBCartItem(c entity.CartItem) CartItem {
	return CartItem{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: c.DeletedAt,
		},
		CartID:      c.CartID,
		ProductName: c.ProductName,
		Quantity:    c.Quantity,
		Price:       c.Price,
	}

}
