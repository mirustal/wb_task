package database

import (
	"context"
	"database/sql"
	"fmt"
	"task-l0/pkg/configs"
)


type Client interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase (cfg *configs.ConfigPostgressDB) (*Database, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Name, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (db *Database) CloseDb()  {
	db.db.Close()
}

func (db *Database) GetDB() *sql.DB {
	return db.db
}

func (db *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.db.BeginTx(ctx, opts)
}