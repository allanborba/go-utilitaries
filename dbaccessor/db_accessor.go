package dbaccessor

import (
	"database/sql"

	"github.com/allanborba/go-utilitaries/postgres"
)

type DBAccessor struct {
	db *sql.DB
	tx *sql.Tx
}

type SQLExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func InitializeDbAccessor(db *sql.DB) *DBAccessor {
	return &DBAccessor{db, nil}
}

func InitializePostgresAccessor() *DBAccessor {
	return InitializeDbAccessor(postgres.ConnectPostgres())
}

func (dbAccessor *DBAccessor) Exec(query string, args ...interface{}) sql.Result {
	result, err := dbAccessor.getExecutor().Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return result
}

func (dbAccessor *DBAccessor) QueryRow(query string, args ...interface{}) *sql.Row {
	return dbAccessor.getExecutor().QueryRow(query, args...)
}

func (dbAccessor *DBAccessor) QueryRows(query string, args ...interface{}) *sql.Rows {
	rows, err := dbAccessor.getExecutor().Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (dbAccessor *DBAccessor) BeginTransaction() {
	tx, err := dbAccessor.db.Begin()
	if err != nil {
		panic(err)
	}
	dbAccessor.tx = tx
}

func (dbAccessor *DBAccessor) CommitTransaction() {
	err := dbAccessor.tx.Commit()
	if err != nil {
		panic(err)
	}
	dbAccessor.tx = nil
}

func (dbAccessor *DBAccessor) RollbackTransaction() {
	if dbAccessor.tx == nil {
		return
	}

	if rec := recover(); rec != nil {
		dbAccessor.tx.Rollback()
		dbAccessor.tx = nil
		panic(rec)
	}
}

func (dbAccessor *DBAccessor) Close() {
	dbAccessor.db.Close()
}

func (dbAccessor *DBAccessor) getExecutor() SQLExecutor {
	if dbAccessor.tx != nil {
		return dbAccessor.tx
	}
	return dbAccessor.db
}
