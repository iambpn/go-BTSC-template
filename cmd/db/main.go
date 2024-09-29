package main

import (
	"database/sql"
	"github.com/iambpn/go-http-template/migrations"
	"os"

	migratorCli "github.com/iambpn/bun-migrator-cli"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/migrate"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbUrl := os.Getenv("DB_URL")
	driverName := os.Getenv("DB_DRIVER")

	sqlDb, err := sql.Open(driverName, dbUrl)

	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqlDb, sqlitedialect.New())
	migratorCli.InitCli(migrate.NewMigrator(db, migrations.Migrations), os.Args)
}
