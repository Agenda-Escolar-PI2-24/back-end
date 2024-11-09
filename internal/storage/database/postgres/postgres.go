package postgres

import (
	"context"
	"log"

	"agenda-escolar/internal/config"

	_ "github.com/lib/pq"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func GetDB() *ksql.DB {
	ctx := context.Background()
	databaseURI := config.POSTGRES_URL_CONN
	dbConnect, err := kpgx.New(ctx, databaseURI, ksql.Config{})
	if err != nil {
		log.Panic(err)
	}
	dbConnect.Exec(ctx, "set enable_seqscan = off;")

	return &dbConnect
}
