package main

import (
	"ctf-challenge/internal/app"
	"ctf-challenge/internal/services"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func initDb() *services.DbService {
	connStr := "postgres://username:password@database:5432/shortlinksDb"
	return services.NewDbService(connStr)
}

func main() {
	err := godotenv.Load("/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbService := initDb()
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	zap.L().Info("Starting server...")
	server := app.SetupRouter(dbService)
	server.Run(":5001")
}
