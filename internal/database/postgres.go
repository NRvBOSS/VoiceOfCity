package database

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/NRvBOSS/citysesi/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	query := u.Query()
	query.Set("sslmode", cfg.DBSSLMode)
	u.RawQuery = query.Encode()

	pool, err := pgxpool.New(ctx, u.String())
	if err != nil {
		return nil, fmt.Errorf("PostgreSQL bağlantısı yaradıla bilmədi: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("PostgreSQL əlçatan deyil: %w", err)
	}

	return pool, nil
}
