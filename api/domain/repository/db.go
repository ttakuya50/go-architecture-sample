package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=$GOFILE -destination=db_mock.go -package=$GOPACKAGE -self_package=github.com/ttakuya50/go-architecture-sample/api/domain/$GOPACKAGE

type DB interface {
	Beginner
	ContextExecutor
}

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
	ContextExecutor
}

type ContextExecutor interface {
	Executor

	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type db struct {
	conn *sql.DB
}

func NewDB(conn *sql.DB) DB {
	boil.SetDB(conn)
	boil.DebugMode = true

	return &db{
		conn: conn,
	}
}

func (d *db) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	return d.conn.BeginTx(ctx, opts)
}

func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.conn.Exec(query, args...)
}

func (d *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.conn.Query(query, args...)
}

func (d *db) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.conn.QueryRow(query, args...)
}

func (d *db) ExecContext(ctx context.Context, s string, i ...interface{}) (sql.Result, error) {
	return d.conn.ExecContext(ctx, s, i...)
}

func (d *db) QueryContext(ctx context.Context, s string, i ...interface{}) (*sql.Rows, error) {
	return d.conn.QueryContext(ctx, s, i...)
}

func (d *db) QueryRowContext(ctx context.Context, s string, i ...interface{}) *sql.Row {
	return d.conn.QueryRowContext(ctx, s, i...)
}
