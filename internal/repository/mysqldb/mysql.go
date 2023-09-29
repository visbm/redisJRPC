package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
)

type MysqlRepository struct {
	DB     *sql.DB
	Logger logger.Logger
}

func NewMysqlDB(cfg config.MySQLStorage, logger logger.Logger) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode))
	if err != nil {
		logger.Panicf("Database open error:%s", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Errorf("DB ping error:%s", err)
		return nil, err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS article(
		id INTEGER PRIMARY KEY,		
		url TEXT NOT NULL,
		topic TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		logger.Errorf("DB Prepare error:%s", err)
		return nil, fmt.Errorf("%s",  err)
	}

	_, err = stmt.Exec()
	if err != nil {
		logger.Errorf("DB Exec error:%s", err)
		return nil, fmt.Errorf("%s", err)
	}

	return &MysqlRepository{
		DB:     db,
		Logger: logger,
	}, nil
}

func (m MysqlRepository) Close() error {
	return m.DB.Close()
}
