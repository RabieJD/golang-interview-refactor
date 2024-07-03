package calculator

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"interview/pkg/controllers"
	"interview/pkg/entity"
)

type TTAddItem struct {
	Name           string
	ItemForm       controllers.CartItemForm
	SessionID      string
	SetupMocks     func(cartRepo *CartRepoMock, priceRepo *PriceRepoMock)
	ExpectedStatus int
	ExpectedErr    error
}

func GenerateTTAddItem() []TTAddItem {
	validGetCartBySessionID := func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
		return &entity.CartEntity{
			Model:     gorm.Model{ID: 1},
			SessionID: SessionID,
			Status:    entity.CartOpen,
		}, nil
	}
	validGetItemByCartIDAndName := func(ctx context.Context, cartID uint, Name string) (*entity.CartItem, error) {
		return &entity.CartItem{
			Model:       gorm.Model{ID: 1},
			CartID:      cartID,
			ProductName: Name,
			Quantity:    1,
			Price:       10.0,
		}, nil
	}
	return []TTAddItem{
		{
			Name: "New Cart and Item",
			ItemForm: controllers.CartItemForm{
				Product:  "TestItem",
				Quantity: "2",
			},
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock, priceRepo *PriceRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return nil, gorm.ErrRecordNotFound
				}
				cartRepo.AddCartFunc = func(ctx context.Context, cart *entity.CartEntity) error {
					return nil
				}
				cartRepo.UpdateItemFunc = func(ctx context.Context, item *entity.CartItem) error {
					return nil
				}
				priceRepo.GetPriceFunc = func(identifier string) (float64, error) {
					return 10.0, nil
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    nil,
		},
		{
			Name: "Existing Cart and Item",
			ItemForm: controllers.CartItemForm{
				Product:  "TestItem",
				Quantity: "2",
			},
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock, priceRepo *PriceRepoMock) {
				cartRepo.GetCartBySessionIDFunc = validGetCartBySessionID
				cartRepo.GetItemByCartIDAndNameFunc = validGetItemByCartIDAndName
				cartRepo.UpdateItemFunc = func(ctx context.Context, item *entity.CartItem) error {
					return nil
				}
				priceRepo.GetPriceFunc = func(identifier string) (float64, error) {
					return 10.0, nil
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    nil,
		},
		{
			Name: "Invalid Item Name",
			ItemForm: controllers.CartItemForm{
				Product:  "invalid",
				Quantity: "2",
			},
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock, priceRepo *PriceRepoMock) {
				priceRepo.GetPriceFunc = func(identifier string) (float64, error) {
					return 0, errors.New("invalid item name")
				}
				cartRepo.GetCartBySessionIDFunc = validGetCartBySessionID
				cartRepo.GetItemByCartIDAndNameFunc = validGetItemByCartIDAndName
			},
			ExpectedStatus: 302,
			ExpectedErr:    errors.New("invalid item name"),
		},
		{
			Name: "Invalid Quantity",
			ItemForm: controllers.CartItemForm{
				Product:  "TestItem",
				Quantity: "invalid",
			},
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock, priceRepo *PriceRepoMock) {
				cartRepo.GetCartBySessionIDFunc = validGetCartBySessionID
				cartRepo.GetItemByCartIDAndNameFunc = validGetItemByCartIDAndName
				priceRepo.GetPriceFunc = func(identifier string) (float64, error) {
					return 10.0, nil
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    errors.New("invalid quantity"),
		},
	}
}

type TTGetCartData struct {
	Name         string
	SessionID    string
	SetupMocks   func(cartRepo *CartRepoMock)
	ExpectedData []controllers.CartData
	ExpectedErr  error
}

func GenerateTTGetCartData() []TTGetCartData {
	return []TTGetCartData{
		{
			Name:      "Get Cart Data Success",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return &entity.CartEntity{
						Model:     gorm.Model{ID: 1},
						SessionID: SessionID,
						Status:    entity.CartOpen,
					}, nil
				}
				cartRepo.GetItemsByCartIDFunc = func(ctx context.Context, cartID uint) ([]*entity.CartItem, error) {
					return []*entity.CartItem{
						{
							Model:       gorm.Model{ID: 1},
							CartID:      cartID,
							ProductName: "TestItem",
							Quantity:    1,
							Price:       10.0,
						},
					}, nil
				}
			},
			ExpectedData: []controllers.CartData{
				{
					controllers.CartDataIDKey:       uint(1),
					controllers.CartDataQuantityKey: 1,
					controllers.CartDataPriceKey:    10.0,
					controllers.CartDataProductKey:  "TestItem",
				},
			},
			ExpectedErr: nil,
		},
		{
			Name:      "Cart Not Found",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return nil, errors.New("cart not found")
				}
			},
			ExpectedData: nil,
			ExpectedErr:  errors.New("cart not found"),
		},
		{
			Name:      "Error Getting Items",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return nil, errors.New("cart not found")
				}
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return &entity.CartEntity{
						Model:     gorm.Model{ID: 1},
						SessionID: SessionID,
						Status:    entity.CartOpen,
					}, nil
				}
				cartRepo.GetItemsByCartIDFunc = func(ctx context.Context, cartID uint) ([]*entity.CartItem, error) {
					return nil, errors.New("error getting items")
				}
			},
			ExpectedData: nil,
			ExpectedErr:  errors.New("error getting items"),
		},
	}
}

type TTDeleteItem struct {
	Name           string
	ItemID         string
	SessionID      string
	SetupMocks     func(cartRepo *CartRepoMock)
	ExpectedStatus int
	ExpectedErr    error
}

func GenerateTTDeleteItem() []TTDeleteItem {
	return []TTDeleteItem{
		{
			Name:      "Delete Item Success",
			ItemID:    "1",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return &entity.CartEntity{
						Model:     gorm.Model{ID: 1},
						SessionID: SessionID,
						Status:    entity.CartOpen,
					}, nil
				}
				cartRepo.DeleteItemFunc = func(ctx context.Context, ItemID uint) error {
					return nil
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    nil,
		},
		{
			Name:      "Cart Not Found",
			ItemID:    "1",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return nil, errors.New("cart not found")
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    errors.New("cart not found"),
		},
		{
			Name:      "Invalid Item ID",
			ItemID:    "invalid",
			SessionID: "session123",
			SetupMocks: func(cartRepo *CartRepoMock) {
				cartRepo.GetCartBySessionIDFunc = func(ctx context.Context, SessionID string) (*entity.CartEntity, error) {
					return &entity.CartEntity{
						Model:     gorm.Model{ID: 1},
						SessionID: SessionID,
						Status:    entity.CartOpen,
					}, nil
				}
			},
			ExpectedStatus: 302,
			ExpectedErr:    errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},
	}
}
