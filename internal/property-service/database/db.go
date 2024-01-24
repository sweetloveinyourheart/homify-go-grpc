package database

import (
	"fmt"
	"homify-go-grpc/internal/property-service/configs"
	"homify-go-grpc/internal/property-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitPostgresConnection() *gorm.DB {
	configurations := configs.GetConfig()

	host := configurations.PostgresHost
	port := configurations.PostgresPort
	user := configurations.PostgresUser
	password := configurations.PostgresPassword
	dbName := configurations.PostgresDatabase

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	MigrateDatabase(db)

	database = db
	return db
}

func MigrateDatabase(db *gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(
		&models.Amenity{},
		&models.Category{},
		&models.Destination{},
		&models.Property{},
	)

	if err != nil {
		panic(err)
	}
}

// Using this function to get a connection
func GetDB() *gorm.DB {
	return database
}
