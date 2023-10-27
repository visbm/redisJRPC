package pgsqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
)

type PgSqlDB struct {
	DB     *sql.DB
	logger logger.Logger
}

func NewPgSqlDB() *PgSqlDB {
	return &PgSqlDB{}
}

func (m *PgSqlDB) Open(cfg config.Config, logger logger.Logger)  error {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PgSQL.Host, cfg.PgSQL.Port, cfg.PgSQL.Username, cfg.PgSQL.Password, cfg.PgSQL.DBName, cfg.PgSQL.SSLMode)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		logger.Panicf("Database open error:%s", err)
		return  err
	}
	err = db.Ping()
	if err != nil {
		logger.Errorf("DB ping error:%s", err)
		return  err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS articles(
		id TEXT PRIMARY KEY,		
		url TEXT NOT NULL,
		title TEXT NOT NULL);
	`)
	if err != nil {
		logger.Errorf("DB Prepare error:%s", err)
		return  fmt.Errorf("%s", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		logger.Errorf("DB Exec error:%s", err)
		return  fmt.Errorf("%s", err)
	}

	m.DB = db
	m.logger = logger
	return nil
}

func (m *PgSqlDB) Close() error {
	return m.DB.Close()
}
