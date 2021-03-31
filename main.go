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

func main() *gorm.DB {
	dsn := fmt.Sprintf(
		GetEnvWithKey("DB_USER"),
		":",
		GetEnvWithKey("DB_PASSWORD"),
		"@tcp(",
		GetEnvWithKey("DB_HOST"),
		":",
		GetEnvWithKey("DB_PORT"),
		"/",
		GetEnvWithKey("DB_DATABASE"),
		"?charset=utf8mb4&parseTime=True&loc=Local",
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	}

	return DB
}