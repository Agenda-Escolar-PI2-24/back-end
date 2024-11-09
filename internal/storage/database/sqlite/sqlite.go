package sqlite

import (
	"context"
	"log"

	"agenda-escolar/internal/config"

	_ "github.com/lib/pq"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/ksqlite3"
)

func GetDB() *ksql.DB {
	ctx := context.Background()
	databaseURI := config.SQLITE_URL_CONN
	dbConnect, err := ksqlite3.New(ctx, databaseURI, ksql.Config{})
	if err != nil {
		log.Panic(err)
	}
	dbConnect.Exec(ctx, "set enable_seqscan = off;")

	return &dbConnect
}
