package sql_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"interview/pkg/entity"
	. "interview/pkg/infra/database/cart/sql"
	td "interview/pkg/testdata/infra/database/cart/sql"
	"testing"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return gormDB, mock, err
}

func runTestCases(t *testing.T, cases []td.TestCase, testFunc func(tc td.TestCase)) {
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			testFunc(tc)
		})
	}
}

func TestRepo_AddCart(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateADDCartTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		err := repo.AddCart(context.TODO(), tc.Args.(*entity.CartEntity))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_UpdateCart(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateUpdateCartTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		err := repo.UpdateCart(context.TODO(), tc.Args.(*entity.CartEntity))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_AddItem(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateAddItemTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		err := repo.AddItem(context.TODO(), tc.Args.(*entity.CartItem))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_UpdateItem(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateUpdateItemTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		err := repo.UpdateItem(context.TODO(), tc.Args.(*entity.CartItem))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_GetCartBySessionID(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateGetCartBySessionIDTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		_, err := repo.GetCartBySessionID(context.TODO(), tc.Args.(string))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_GetItemByCartIDAndName(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateGetItemByCartIDAndNameTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		_, err := repo.GetItemByCartIDAndName(context.TODO(), tc.Args.([]interface{})[0].(uint), tc.Args.([]interface{})[1].(string))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_GetItemsByCartID(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateGetItemsByCartIDTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		_, err := repo.GetItemsByCartID(context.TODO(), tc.Args.(uint))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepo_DeleteItem(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo, _ := NewRepo(db)

	testCases := td.GenerateDeleteItemTT()

	runTestCases(t, testCases, func(tc td.TestCase) {
		tc.Mock(mock)
		err := repo.DeleteItem(context.TODO(), tc.Args.(uint))
		if tc.HasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
