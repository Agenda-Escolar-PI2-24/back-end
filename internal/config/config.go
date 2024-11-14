package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	SQLITE_URL_CONN   string = ""
	POSTGRES_URL_CONN string = "host=localhost port=5432 user=postgres password=postgres dbname=schedule sslmode=disable"
)

func init() {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	SQLITE_URL_CONN = fmt.Sprintf("%s\\database.db", basepath)
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	iPort, _ := strconv.Atoi(DB_PORT)
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_HOST != "" && iPort != 0 {
		POSTGRES_URL_CONN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, iPort, DB_USER, DB_PASSWORD, DB_NAME)
	}
}
