package connect_mysql

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

//ConnectToDb: connect to mysql with preloaded .env values
func ConnectToDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s%s",
		GetEnvWithKey("DB_USER"),
		":",
		GetEnvWithKey("DB_PASSWORD"),
		"@tcp(",
		GetEnvWithKey("DB_HOST"),
		":",
		GetEnvWithKey("DB_PORT"),
		")/",
		GetEnvWithKey("DB_DATABASE"),
		"?charset=utf8mb4&parseTime=True&loc=Local",
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return DB, err
}