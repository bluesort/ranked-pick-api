package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-choice/api"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting ranked-choice API")

	cfg := api.ParseConfig()

	// TODO: Move DB URL into env var, use postgres role, and enable SSL
	dbMigrate, err := migrate.New(
		"file://db/migrations",
		"sqlite3://db/sqlite3.db",
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
		log.Printf("Database is dirty, forcing migration version %d", dbForceVersion)
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

	router := chi.NewRouter()
	api.AddMiddleware(router)
	api.AddRoutes(router)

	log.Printf("Router listening on port %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		log.Println(err)
	}
}
