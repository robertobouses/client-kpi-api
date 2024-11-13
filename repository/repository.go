package repository

import "database/sql"

//go:embed sql/get_lineup.sql
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
