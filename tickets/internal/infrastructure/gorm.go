package infrastructure

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(config *viper.Viper) *gorm.DB {
	user := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	dbname := config.GetString("DB_NAME")
	port := config.GetInt("DB_PORT")
	idleConns := config.GetInt("POOL_IDLE")
	maxConns := config.GetInt("POOL_MAX")
	lifetime := config.GetDuration("POOL_LIFETIME") * time.Second

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	var db *gorm.DB
	var err error
	maxRetries := 10
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err == nil {
			fmt.Printf("Connected to MySQL database!!!\n")
			break
		}

		fmt.Printf("Error connecting to database: %v\n", err.Error())
		time.Sleep(retryInterval)
	}

	if err != nil {
		panic(fmt.Errorf("error connecting database : %+v", err.Error()))
	}

	connection, err := db.DB()
	if err != nil {
		panic(err)
	}

	connection.SetMaxIdleConns(idleConns)
	connection.SetMaxOpenConns(maxConns)
	connection.SetConnMaxLifetime(lifetime)

	return db
}
