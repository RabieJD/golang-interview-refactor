package sql_test

import (
	"github.com/stretchr/testify/assert"
	. "interview/pkg/infra/database/cart/sql"
	"interview/pkg/testdata/infra/database/cart/sql"
	"testing"
)

func TestCart_ToCartEntity(t *testing.T) {
	for _, tc := range sql.GenerateTTToCartEntity() {
		result := tc.Input.ToCartEntity()
		assert.Equal(t, tc.Expected, result)
	}
}

func TestCartEntityToDBCart(t *testing.T) {
	for _, tc := range sql.GenerateTTCartEntityToDBCart() {
		result := CartEntityToDBCart(tc.Input)
		assert.Equal(t, tc.Expected, result)
	}
}

func TestCart_ToEntityCartInfo2(t *testing.T) {
	for _, tc := range sql.GenerateTTToEntityCartItem() {
		result := tc.Input.ToEntityCartItem()
		assert.Equal(t, tc.Expected, result)
	}
}

func TestCartItemToDBCartItem(t *testing.T) {
	for _, tc := range sql.GenerateTTCartItemToDBCartItem() {
		result := CartItemToDBCartItem(tc.Input)
		assert.Equal(t, tc.Expected, result)
	}
}
