package database

import (
	"agenda-escolar/internal/storage/database/postgres"
	"agenda-escolar/internal/storage/database/sqlite"
	"os"

	"github.com/vingarcia/ksql"
)

func GetDB() *ksql.DB {
	if os.Getenv("production") == "1" {
		return postgres.GetDB()
	}

	return sqlite.GetDB()
}
