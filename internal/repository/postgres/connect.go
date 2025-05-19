package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/ypxd99/yandex-diplom-56/util"
)

type PostgresRepo struct {
	db *bun.DB
}

func Connect(ctx context.Context) (*PostgresRepo, error) {
	cfg := util.GetConfig().Postgres

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.ConnString)))
	db := bun.NewDB(sqlDB, pgdialect.New())

	if cfg.Trace {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	db.DB.SetMaxOpenConns(cfg.MaxConn)
	db.DB.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifeTime) * time.Second)
	//db.Exec(`SET search_path TO gophermart, public;`)

	return &PostgresRepo{db: db}, db.PingContext(ctx)
}

func (p *PostgresRepo) Close() error {
	return p.db.Close()
}

func (p *PostgresRepo) Status(ctx context.Context) (bool, error) {
	err := p.db.PingContext(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
