package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

var (
	ConnectionError = errors.New("Ошибка подключения к БД")
	PingError       = errors.New("База данных недоступна")
)

func NewStorage(ctx context.Context, connString string) (*Storage, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, ConnectionError
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, PingError
	}
	return &Storage{pool: pool}, nil
}

func (s *Storage) Close() {
	if s != nil {
		s.pool.Close()
	}
}
