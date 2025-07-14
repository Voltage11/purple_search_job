package database

import (
	"context"
	"lesson/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbPool, err := pgxpool.New(context.Background(), config.Url)
	if (err != nil) {
		logger.Error().Msg("Не удалось подключиться к БД")
		panic(err)
	}
	err = dbPool.Ping(context.Background())
	if (err != nil) {
		logger.Error().Msg("Нет пинга к БД")
		panic(err)
	}


	logger.Info().Msg("Подключились к БД")
	return  dbPool

}