package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

/**
 * Массив настроек подключения к базе данных
 *
 */
var dbConfig = mysql.Config{
	User:                 "",
	Passwd:               "",
	Net:                  "",
	Addr:                 "",
	DBName:               "",
	AllowNativePasswords: true, // установить false если пароль (DB_PASS в .env) не пустой
}

/**
 * Инициализация подключения к базе данных
 *
 */
func initDateBase() (db *sql.DB) {
	errenv := godotenv.Load(".env")
	if errenv != nil {
		fmt.Sprintln("Error loading .env file")
	}
	dbConfig.User = os.Getenv("DB_USER")
	dbConfig.Passwd = os.Getenv("DB_PASS")
	dbConfig.Net = os.Getenv("DB_NET")
	dbConfig.Addr = os.Getenv("DB_HOST")
	dbConfig.DBName = os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
