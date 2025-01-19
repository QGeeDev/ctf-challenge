package main

import (
	"ctf-challenge/internal/app"
	"ctf-challenge/internal/services"

	"go.uber.org/zap"
)

func initDb() *services.DbService {
	connStr := "postgres://username:password@database:5432/shortlinksDb"
	return services.NewDbService(connStr)
}

func main() {
	dbService := initDb()
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	zap.L().Info("Starting server...")
	server := app.SetupRouter(dbService)
	server.Run(":5001")
}
