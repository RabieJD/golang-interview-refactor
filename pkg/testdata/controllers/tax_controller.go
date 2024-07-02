package controllers

import (
	"context"
	"errors"
	"interview/pkg/controllers"
	"net/http"
)

type TTAddItem struct {
	Name     string
	Setup    func(*CalculatorHandlerMock)
	Cookie   *http.Cookie
	FormData map[string]string
	Expected int
}

func GenerateTTAddItem() []TTAddItem {
	return []TTAddItem{
		{
			Name: "successful add item",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.AddItemFunc = func(ctx context.Context, item controllers.CartItemForm, sessionID string) (int, error) {
					return http.StatusFound, nil
				}
			},
			Cookie:   &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			FormData: map[string]string{"product": "item1", "quantity": "1"},
			Expected: http.StatusFound,
		},
		{
			Name: "missing session ID",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.AddItemFunc = func(ctx context.Context, item controllers.CartItemForm, sessionID string) (int, error) {
					return http.StatusFound, nil
				}
			},
			Cookie:   nil,
			FormData: map[string]string{"product": "item1", "quantity": "1"},
			Expected: http.StatusFound,
		},
		{
			Name: "error adding item",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.AddItemFunc = func(ctx context.Context, item controllers.CartItemForm, sessionID string) (int, error) {
					return http.StatusFound, errors.New("add item error")
				}
			},
			Cookie:   &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			FormData: map[string]string{"product": "item1", "quantity": "1"},
			Expected: http.StatusFound,
		},
	}
}

type TTTShowAddItemForm struct {
	Name        string
	Setup       func(*CalculatorHandlerMock)
	QueryParams map[string]string
	Cookie      *http.Cookie
	Expected    int
}

func GenerateTTTShowAddItemForm() []TTTShowAddItemForm {
	return []TTTShowAddItemForm{
		{
			Name: "successful form display",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.GetCartDataFunc = func(ctx context.Context, sessionID string) ([]controllers.CartData, error) {
					return []controllers.CartData{}, nil
				}
			},
			Cookie:   &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			Expected: http.StatusOK,
		},
		{
			Name: "missing Cookie",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.GetCartDataFunc = func(ctx context.Context, sessionID string) ([]controllers.CartData, error) {
					return []controllers.CartData{}, nil
				}
			},
			Cookie:   nil,
			Expected: http.StatusOK,
		},
		{
			Name: "error fetching cart data",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.GetCartDataFunc = func(ctx context.Context, sessionID string) ([]controllers.CartData, error) {
					return nil, errors.New("cart error")
				}
			},
			Cookie:   &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			Expected: 200,
		},
	}
}

type TTDeleteItem struct {
	Name        string
	Setup       func(*CalculatorHandlerMock)
	QueryParams map[string]string
	Cookie      *http.Cookie
	Expected    int
}

func GenerateTTDeleteItem() []TTDeleteItem {
	return []TTDeleteItem{
		{
			Name: "successful delete item",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.DeleteItemFromCartFunc = func(ctx context.Context, itemID string, sessionID string) (int, error) {
					return http.StatusFound, nil
				}
			},
			Cookie:      &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			QueryParams: map[string]string{"cart_item_id": "item1"},
			Expected:    http.StatusFound,
		},
		{
			Name: "missing session ID",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.DeleteItemFromCartFunc = func(ctx context.Context, itemID string, sessionID string) (int, error) {
					return http.StatusFound, nil
				}
			},
			Cookie:      nil,
			QueryParams: map[string]string{"cart_item_id": "item1"},
			Expected:    http.StatusFound,
		},
		{
			Name: "error deleting item",
			Setup: func(mock *CalculatorHandlerMock) {
				mock.DeleteItemFromCartFunc = func(ctx context.Context, itemID string, sessionID string) (int, error) {
					return http.StatusFound, errors.New("delete item error")
				}
			},
			Cookie:      &http.Cookie{Name: "ice_session_id", Value: "test-session-id"},
			QueryParams: map[string]string{"cart_item_id": "item1"},
			Expected:    http.StatusFound,
		},
	}
}
