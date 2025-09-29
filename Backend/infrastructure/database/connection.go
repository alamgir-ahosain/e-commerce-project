package database

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DbConfig) string {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Name)
	//fmt.Println("user:", cnf.User, "pass: ", cnf.Password, " host:", cnf.Host, "name: ", cnf.Name)
	if !cnf.EnableSSLMode {
		connectionStr += " sslmode=disable"
	}

	//fmt.Println("conStr:", connectionStr)
	return connectionStr
}

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbConnection, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil

}
