package services

import (
	"context"
	"crypto/rand"
	"ctf-challenge/internal/dto"
	"encoding/base64"

	"github.com/jackc/pgx/v5"
)

type ShortlinkService struct {
	DbConn pgx.Conn
}

func NewShortlinkService(dbConn pgx.Conn) *ShortlinkService {
	return &ShortlinkService{
		DbConn: dbConn,
	}
}

func (s *ShortlinkService) GetFullUrlByShortlinkSlug(slug string) (string, error) {
	var fullLink string
	err := s.DbConn.QueryRow(context.Background(), "SELECT full_link FROM shortlinks WHERE slug=$1", slug).Scan(&fullLink)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return fullLink, nil
}

func (s *ShortlinkService) CreateShortlink(targetUrl string) (*dto.ShortlinkDbObject, error) {
	result, err := s.generateBase64String(8)
	if err != nil {
		return nil, err
	}

	ShortlinkRecord := &dto.ShortlinkDbObject{
		Slug:     result,
		FullLink: targetUrl,
	}
	var id int
	err = s.DbConn.QueryRow(context.Background(), "INSERT INTO shortlinks (slug, full_link) VALUES ($1, $2) returning Id", ShortlinkRecord.Slug, ShortlinkRecord.FullLink).Scan(&id)
	if err != nil {
		return nil, err
	}

	return ShortlinkRecord, err
}

func (s *ShortlinkService) generateBase64String(length int) (string, error) {
	byteLength := (length * 3) / 4

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)
	return encoded[:length], nil
}

func (s *ShortlinkService) GetQrCodePath(shortlinkSlug string) (any, error) {
	var path string
	err := s.DbConn.QueryRow(context.Background(), "SELECT q.image_path FROM shortlinks s WHERE s.slug=$1 INNER JOIN qr_images q on q.id = s.qr_images_id_fk", shortlinkSlug).Scan(&path)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return path, nil
}
