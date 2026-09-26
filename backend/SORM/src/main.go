package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/SuliR123/Journal-Time/SORM/src/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Printf("Current working directory %s\n", dir)

	// TODO: Setup where the user's configuration file will be
	err = godotenv.Load(filepath.Join(dir, "SORM", "src", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err)
	}

	dbUrl := os.Getenv("DB_CONNECTION_STRING")

	dbpool, err := pgxpool.New(context.Background(), dbUrl)

	modelsDirPath := dir + os.Getenv("MODEL_DIRECTORY")
	fmt.Printf("Path to the model directory: %s \n", modelsDirPath)
	// TODO: Allow the user to configure environments with set path values and what not (for dev db and prod db ts)
	migrationCommands := migrations.New(dbpool, modelsDirPath)
	fmt.Printf("Migration Commands: %s", migrationCommands.String())
}
