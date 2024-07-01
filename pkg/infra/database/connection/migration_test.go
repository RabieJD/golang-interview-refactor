package connection_test

import (
	. "interview/pkg/infra/database/connection"
	td "interview/pkg/testdata/infra/database/connection"
	"testing"
)

func TestDBMigration_MigrateDBs(t *testing.T) {
	for _, tt := range td.GenerateTTMigrateDB() {
		t.Run(tt.Name, func(t *testing.T) {
			dbMigration := NewDBMigration(tt.Migrations)
			err := dbMigration.MigrateDBs()
			if (err != nil) != tt.HasError {
				t.Errorf("MigrateDBs() error = %v, wantErr %v", err, tt.HasError)
			}
		})
	}
}

func TestDBMigration_AddMigrations(t *testing.T) {
	for _, tt := range td.GenerateTTAddMigrations() {
		t.Run(tt.Name, func(t *testing.T) {
			dbMigration := NewDBMigration(tt.Initial)
			dbMigration.Add(tt.ToAdd...)
			// we can not access the length so we will count using migrate instead
			_ = dbMigration.MigrateDBs()
			if int(*tt.Length) != tt.ExpectedLength {
				t.Errorf("Add() got %d migrations, want %d", int(*tt.Length), tt.ExpectedLength)
			}
		})
	}
}
