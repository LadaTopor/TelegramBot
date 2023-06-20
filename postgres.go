package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func PostgresConnection() (*sql.DB, error) {
	// Строка подключения к базе данных PostgreSQL
	connStr := "user=postgres dbname=study-schedule password=asdfg host=localhost sslmode=disable"

	// Открытие соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, err
}
