package database

import (
	"fmt"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewClient(logger logger.LoggerInterface) (*sqlx.DB, error) {
	dbDriver := viper.GetString("DB_DRIVER")

	if dbDriver != "postgres" && dbDriver != "pgx" {
		logger.Error("sqlx only supports PostgreSQL", zap.String("DB_DRIVER", dbDriver))
		return nil, fmt.Errorf("sqlx only supports PostgreSQL, got: %s", dbDriver)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
	)

	maxOpenConns := viper.GetInt("DB_MAX_OPEN_CONNS")
	if maxOpenConns <= 0 {
		maxOpenConns = 100
	}

	minIdleConns := viper.GetInt("DB_MIN_IDLE_CONNS")
	if minIdleConns <= 0 {
		minIdleConns = 50
	}
	if minIdleConns > maxOpenConns {
		minIdleConns = maxOpenConns
	}

	connMaxLifetime := viper.GetDuration("DB_CONN_MAX_LIFETIME")
	if connMaxLifetime == 0 {
		connMaxLifetime = time.Hour
	}

	connMaxIdleTime := viper.GetDuration("DB_CONN_MAX_IDLE_TIME")
	if connMaxIdleTime == 0 {
		connMaxIdleTime = 30 * time.Minute
	}

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(minIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	logger.Debug("Database connection pool established successfully",
		zap.String("DB_DRIVER", "pgx"),
		zap.Int("MaxOpenConns", maxOpenConns),
		zap.Int("MaxIdleConns", minIdleConns),
		zap.Duration("ConnMaxLifetime", connMaxLifetime),
		zap.Duration("ConnMaxIdleTime", connMaxIdleTime),
	)

	return db, nil
}
