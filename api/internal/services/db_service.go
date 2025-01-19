package services

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type DbService struct {
	DbConn pgx.Conn
}

func NewDbService(connString string) *DbService {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		os.Exit(1)
	}
	return &DbService{
		DbConn: *conn,
	}
}
