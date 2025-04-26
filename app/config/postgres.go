package config

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"false"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"false"`
	User     string `envconfig:"POSTGRES_USER" required:"false"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"false"`
	Dbname   string `envconfig:"POSTGRES_DATABASE" required:"false"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" required:"false"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" required:"false"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" required:"false"`
}

func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func (p Postgres) DSNSecured() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func OpenPostgresConnection(pg Postgres, secured bool) (*gorm.DB, error) {
	dsn := pg.DSN()
	if secured {
		dsn = pg.DSNSecured()
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(pg.MaxConnectionLifetime)
	sqlDB.SetMaxOpenConns(pg.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(pg.MaxIdleConnection)

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db.Debug(), nil
}

func Ping(ctx context.Context, db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Successfully pinged the database.")
	return nil
}
