package database

import (
	"fmt"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukrainskykirill/pepper/pkg/config"
)

func NewPostgresPool(config *config.Config) *pgxpool.Pool{
	ctx := context.Background()
	db, err := pgxpool.New(
		ctx, 
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
			config.PostgresUser, 
			config.PostgresPassword,
			config.PostgresDBName,
			config.PostgresSSLMode, 
			config.PostgresHost,
			config.PostgresPort,
		),
	)
	if err != nil {
		fmt.Println(err)
	}
	return db
}