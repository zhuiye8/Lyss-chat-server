package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zhuiye8/Lyss-chat-server/pkg/config"
)

// Postgres 表示 PostgreSQL 数据库连�?
type Postgres struct {
	DB *sqlx.DB
}

// NewPostgres 创建一个新�?PostgreSQL 数据库连�?
func NewPostgres(cfg config.DatabaseConfig) (*Postgres, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// 设置连接池参�?
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return &Postgres{DB: db}, nil
}

// Close 关闭数据库连�?
func (p *Postgres) Close() error {
	return p.DB.Close()
}

