package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/stg35/avito_test/internal/config"
)

func NewConn(config *config.DBConfig) (*pg.DB, error) {
	opt, err := pg.ParseURL(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName),
	)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
