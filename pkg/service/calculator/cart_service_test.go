package calculator

import (
	"context"
	"github.com/stretchr/testify/assert"
	td "interview/pkg/testdata/service/calculator"
	"testing"
)

func TestCartService_AddItem(t *testing.T) {
	for _, tt := range td.GenerateTTAddItem() {
		t.Run(tt.Name, func(t *testing.T) {
			cartRepo := &td.CartRepoMock{}
			priceRepo := &td.PriceRepoMock{}
			tt.SetupMocks(cartRepo, priceRepo)
			service := NewCartService(cartRepo, priceRepo)
			status, err := service.AddItem(context.Background(), tt.ItemForm, tt.SessionID)
			assert.Equal(t, tt.ExpectedStatus, status)
			if tt.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.ExpectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCartService_GetCartData(t *testing.T) {
	for _, tt := range td.GenerateTTGetCartData() {
		t.Run(tt.Name, func(t *testing.T) {
			cartRepo := &td.CartRepoMock{}
			tt.SetupMocks(cartRepo)
			service := NewCartService(cartRepo, nil)
			data, err := service.GetCartData(context.Background(), tt.SessionID)
			if tt.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.ExpectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.ExpectedData, data)
			}
		})
	}
}

func TestCartService_DeleteItemFromCart(t *testing.T) {
	for _, tt := range td.GenerateTTDeleteItem() {
		t.Run(tt.Name, func(t *testing.T) {
			cartRepo := &td.CartRepoMock{}
			tt.SetupMocks(cartRepo)
			service := NewCartService(cartRepo, nil)
			status, err := service.DeleteItemFromCart(context.Background(), tt.ItemID, tt.SessionID)
			assert.Equal(t, tt.ExpectedStatus, status)
			if tt.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.ExpectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
