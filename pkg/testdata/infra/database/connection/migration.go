package connection

import (
	"errors"
	"interview/pkg/infra/database/connection"
	"sync/atomic"
)

// MockMigration is a mock implementation of the Migration interface
type MockMigration struct {
	ShouldFail bool
	CalledNB   *int32
}

func (m *MockMigration) Migrate() error {
	atomic.AddInt32(m.CalledNB, 1)
	if m.ShouldFail {
		return errors.New("migration failed")
	}
	return nil
}

type TTMigrateDB struct {
	Name       string
	Migrations []connection.Migration
	HasError   bool
}

func GenerateTTMigrateDB() []TTMigrateDB {
	return []TTMigrateDB{
		{
			Name: "All migrations succeed",
			Migrations: []connection.Migration{
				&MockMigration{ShouldFail: false},
				&MockMigration{ShouldFail: false},
			},
			HasError: false,
		},
		{
			Name: "One migration fails",
			Migrations: []connection.Migration{
				&MockMigration{ShouldFail: false},
				&MockMigration{ShouldFail: true},
			},
			HasError: true,
		},
		{
			Name:       "No migrations",
			Migrations: []connection.Migration{},
			HasError:   false,
		},
	}
}

type TTAddMigrations struct {
	Name           string
	Initial        []connection.Migration
	ToAdd          []connection.Migration
	Length         *int32
	ExpectedLength int
}

func GenerateTTAddMigrations() []TTAddMigrations {
	var length1, length2, length3 int32
	return []TTAddMigrations{
		{
			Name:           "Add single migration",
			Initial:        []connection.Migration{},
			ToAdd:          []connection.Migration{&MockMigration{ShouldFail: false, CalledNB: &length1}},
			Length:         &length1,
			ExpectedLength: 1,
		},
		{
			Name: "Add multiple migrations",
			Initial: []connection.Migration{
				&MockMigration{ShouldFail: false, CalledNB: &length2},
			},
			ToAdd: []connection.Migration{
				&MockMigration{ShouldFail: false, CalledNB: &length2},
				&MockMigration{ShouldFail: true, CalledNB: &length2},
			},
			Length:         &length2,
			ExpectedLength: 3,
		},
		{
			Name:           "Add to empty initial list",
			Initial:        []connection.Migration{},
			ToAdd:          []connection.Migration{},
			Length:         &length3,
			ExpectedLength: 0,
		},
	}
}
