package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const databaseConnectionExtraParameters = "charset=utf8mb4&parseTime=True&loc=Local"

func GetDBConnection(conf *DatabaseConnection) (*gorm.DB, error) {
	// MySQL connection string
	// Update the username, password, host, port, and database name accordingly
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)%s?%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
		databaseConnectionExtraParameters)

	// Open the connection to the database
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
