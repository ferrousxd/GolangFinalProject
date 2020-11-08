package repositories

import (
	"database/sql"
	"sync"
)

type Database interface {
	GetConnection() *sql.DB
}

type singletonDatabase struct {
	connection *sql.DB
}

func (s *singletonDatabase) GetConnection() *sql.DB {
	return s.connection
}

func getConnection() *sql.DB {
	connStr := "user=postgres port=5432 password=140499 dbname=go_project sslmode=disable"

	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return conn
}

var once sync.Once
var instance Database

func GetSingletonDatabase() Database {
	once.Do(func() {
		conn := getConnection()
		db := singletonDatabase{conn}
		instance = &db
	})

	return instance
}