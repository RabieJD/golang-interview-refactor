package connection

type Migration interface {
	Migrate() error
}

type DBMigration struct {
	migrations []Migration
}

func NewDBMigration(migrations []Migration) *DBMigration {
	return &DBMigration{migrations: migrations}
}

func (m *DBMigration) MigrateDBs() error {
	for _, migration := range m.migrations {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}
	return nil
}

func (m *DBMigration) Add(migrations ...Migration) {
	m.migrations = append(m.migrations, migrations...)
}
