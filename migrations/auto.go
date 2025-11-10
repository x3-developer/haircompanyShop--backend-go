package main

import (
	"database/sql"
	"errors"
	_ "flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
	"serv_shop_haircompany/internal/config"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := sql.Open("postgres", persistence.GetDSN(cfg))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run auto.go [up|down]")
		os.Exit(1)
	}
	command := os.Args[1]

	switch command {
	case "up":
		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Error applying migrations up: %v", err)
		}
		fmt.Println("Migrations applied up successfully.")
	case "down":
		err = m.Down()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Error applying migrations down: %v", err)
		}
		fmt.Println("Migrations applied down successfully.")
	default:
		fmt.Println("Unknown command. Use 'up' or 'down'.")
		os.Exit(1)
	}
}
