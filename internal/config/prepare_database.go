package config

import (
	"database/sql"
	"log"

	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PrepareDatabase() {
	dbMigrate, err := migrate.New(
		"file://migrations",
		"sqlite3://sqlite3.db?x-no-tx-wrap=true",
	)
	if err != nil {
		log.Fatal(err)
	}

	dbVersion, dbDirty, err := dbMigrate.Version()
	if err == migrate.ErrNilVersion {
		log.Println("Initializing database")
	} else if err != nil {
		log.Fatal(err)
	}

	if dbDirty {
		dbForceVersion := dbVersion - 1
		log.Printf("Database is dirty, forcing version %d", dbForceVersion)
		err = dbMigrate.Force(int(dbForceVersion))
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Database at version %d\n", dbVersion)

	err = dbMigrate.Up()
	if err == migrate.ErrNoChange {
		log.Println("No migrations to be run")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Migrations run")
	}

	// Init db connection
	Config.Db, err = sql.Open("sqlite3", env.GetString("DB_URL", "sqlite3.db"))
	if err != nil {
		log.Fatal(err)
	}
	Config.Queries = db.New(Config.Db)
}
