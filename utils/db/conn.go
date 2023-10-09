package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *Database

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Conn     *sql.DB
}

// InitDB initializes the database connection
func InitDB(host, port, user, password, database string) {
	if DB != nil {
		return
	}
	DB = &Database{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
	var err error
	DB.Conn, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		panic(err.Error())
	}
	go pingMySql()
	log.Println("connected to database")
}

// pingMySql pings the database every minute to keep the connection alive
func pingMySql() {
	for {
		_, err := DB.Conn.Exec("SELECT 1")
		if err != nil {
			panic(err.Error())
		}
		// sleep for 1 minute
		<-time.After(1 * time.Minute)
	}
}
