package services

import (
	"context"
	"ctf-challenge/internal/dto"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type AdminService struct {
	DbConn pgx.Conn
}

func NewAdminService(dbConn pgx.Conn) *AdminService {
	return &AdminService{
		DbConn: dbConn,
	}
}

func (s *AdminService) AddCustomShortlink(createRequest dto.CreateCustomShortlinkRequest) (string, error) {
	sql_query := fmt.Sprintf("INSERT INTO shortlinks (slug, full_link) VALUES ('%s', '%s') returning id", createRequest.CustomSlug, createRequest.TargetUrl)

	result, err := s.DbConn.Exec(context.Background(), sql_query)

	if err != nil {
		return "", err
	}

	var flag string
	for result.Next() {
		err := result.Scan(&flag)
		if err != nil {
			return "", err
		}
	}

	if flag != "" {
		return flag, nil
	}

	return "", nil
}
