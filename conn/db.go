package conn

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rhasan33/goplate/config"

	// postgres conn
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresClient holds the redis client instance
type PostgresClient struct {
	*gorm.DB
}

// Conn is an instance *gorm.DB
var Conn PostgresClient

// ConnectDB : Setup assigns for postgres db
func ConnectDB() error {
	cfg := config.DB()
	dbSource := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Name,
		cfg.Password,
	)

	c, err := gorm.Open("postgres", dbSource)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("database connection error")
		return err
	}
	fmt.Println("successful db connection")
	Conn = PostgresClient{
		DB: c,
	}
	// need to pass the references of model
	Conn.DB.AutoMigrate()
	return nil
}

// PostgresDB to get db connections
func PostgresDB() PostgresClient {
	return Conn
}
