package config

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	SQLITE_URL_CONN   string = ""
	POSTGRES_URL_CONN string = "host=localhost port=5432 user=postgres password=postgres dbname=agenda_escolar sslmode=disable"
)

func init() {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	SQLITE_URL_CONN = fmt.Sprintf("%s\\database.db", basepath)
}
