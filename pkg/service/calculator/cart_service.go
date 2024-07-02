package calculator

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"interview/pkg/controllers"
	"interview/pkg/entity"
	"strconv"
)

type cartService struct {
	CartRepo
	PriceRepo
}

func NewCartService(cartRepo CartRepo, priceRepo PriceRepo) controllers.CalculatorHandler {
	return &cartService{CartRepo: cartRepo, PriceRepo: priceRepo}
}

func (c *cartService) AddItem(ctx context.Context, itemForm controllers.CartItemForm, sessionID string) (int, error) {
	// retrieve the cart
	cart, err := c.CartRepo.GetCartBySessionID(ctx, sessionID)
	// if cart exist
	if err == nil {
		item, err := c.CartRepo.GetItemByCartIDAndName(ctx, cart.ID, itemForm.Product)
		// if item exists
		if err == nil {
			if err := c.saveItem(ctx, cart.ID, itemForm, item); err != nil {
				return 302, err
			}
			return 302, nil
		}
		// if item doesn't exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := c.saveItem(ctx, cart.ID, itemForm, nil); err != nil {
				return 302, err
			}
			return 302, nil
		}
		return 302, err
	}
	// if cart doesn't exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// save cart
		cartEntity := &entity.CartEntity{
			SessionID: sessionID,
			Status:    entity.CartOpen,
		}
		if err := c.CartRepo.AddCart(ctx, cartEntity); err != nil {
			return 302, err
		}
		// save item
		if err := c.saveItem(ctx, cartEntity.ID, itemForm, nil); err != nil {
			return 302, err
		}
		return 302, nil
	}
	return 302, err
}

func (c *cartService) saveItem(ctx context.Context, cartID uint, itemForm controllers.CartItemForm, oldItem *entity.CartItem) error {
	itemPrice, err := c.PriceRepo.GetPrice(itemForm.Product)
	if err != nil {
		return errors.New("invalid item name")
	}

	quantity, err := strconv.ParseInt(itemForm.Quantity, 10, 0)
	if err != nil {
		return errors.New("invalid quantity")
	}
	itemToSave := entity.CartItem{CartID: cartID}
	if oldItem != nil {
		itemToSave = *oldItem
	}
	itemToSave.Price += float64(quantity) * itemPrice
	itemToSave.Quantity += int(quantity)
	return c.CartRepo.UpdateItem(ctx, &itemToSave)
}

func (c *cartService) GetCartData(ctx context.Context, sessionID string) ([]controllers.CartData, error) {
	// get teh cartEntity
	cartEntity, err := c.CartRepo.GetCartBySessionID(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	// get cart items
	cartItems, err := c.CartRepo.GetItemsByCartID(ctx, cartEntity.ID)
	if err != nil {
		return nil, err
	}
	// populate the result and return it
	var result []controllers.CartData
	for _, cartItem := range cartItems {
		result = append(result, controllers.CartData{
			controllers.CartDataIDKey:       cartItem.ID,
			controllers.CartDataQuantityKey: cartItem.Quantity,
			controllers.CartDataPriceKey:    cartItem.Price,
			controllers.CartDataProductKey:  cartItem.ProductName,
		})
	}
	return result, nil
}

func (c *cartService) DeleteItemFromCart(ctx context.Context, itemID string, sessionID string) (int, error) {
	// check if the cart exists and is open
	if _, err := c.CartRepo.GetCartBySessionID(ctx, sessionID); err != nil {
		return 302, err
	}
	// convert cartID
	cartItemID, err := strconv.Atoi(itemID)
	if err != nil {
		return 302, err
	}
	// delete Item
	err = c.CartRepo.DeleteItem(ctx, uint(cartItemID))
	return 302, err
}
