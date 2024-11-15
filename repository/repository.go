package repository

import (
	"database/sql"
	_ "embed"
)

//go:embed sql/query_all_clients.sql
var queryAllClients string

func NewRepository(db *sql.DB) (*repository, error) {
	queryAllClientsStmt, err := db.Prepare(queryAllClients)
	if err != nil {
		return nil, err
	}
	return &repository{
		db:              db,
		queryAllClients: queryAllClientsStmt,
	}, nil

}

type repository struct {
	db              *sql.DB
	queryAllClients *sql.Stmt
}
