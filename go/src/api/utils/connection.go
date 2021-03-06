package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func dbURL(dbConfig *dBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// GetConnection exports the DB
func GetConnection() *gorm.DB {
	config := &dBConfig{
		Host:     "app_mariadb",
		Port:     3306,
		User:     "dev",
		DBName:   "fullstackAsociacion",
		Password: "passdev",
	}
	db, err := gorm.Open(mysql.Open(dbURL(config)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error BBDD")
	} else {
		log.Printf("I'm connected!")
	}
	return db
}

// GetConnectionTest exports to the test DB
func GetConnectionTest() *gorm.DB {
	config := &dBConfig{
		Host:     "app_mariadb_fake",
		Port:     3306,
		User:     "dev",
		DBName:   "fullstackAsociacion_test",
		Password: "passdev",
	}
	db, err := gorm.Open(mysql.Open(dbURL(config)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("I'm connected!")
	}
	return db
}

/*
docker-compose up -d
docker network create testAPI
docker build -t tests .
docker run -d --net testAPI --name app_mariadb_fake fs5-gestiondesocios-backend_db
docker run --net testAPI --name test_db tests


*/
// // GetConnectionLocalHost exports the DB
// func GetConnectionLocalHost() *gorm.DB {
// 	config := &dBConfig{
// 		Host:     "127.0.0.1",
// 		Port:     3306,
// 		User:     "dev",
// 		DBName:   "fullstackAsociacion",
// 		Password: "passdev",
// 	}
// 	db, err := gorm.Open(mysql.Open(dbURL(config)), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error BBDD")
// 	} else {
// 		log.Printf("I'm connected!")
// 	}
// 	return db
// }
