package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DbConfig) string {
	conString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cnf.DbHost, cnf.DbPort, cnf.DbUsername, cnf.DbPassword, cnf.DbName)
	if !cnf.EnableSSL {
		conString += " sslmode=disable"
	}
	return conString
}

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	return db, nil
}
