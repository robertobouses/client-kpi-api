package repository

import (
	"database/sql"
	_ "embed"
)

//go:embed sql/query_all_clients.sql
var queryAllClients string

//go:embed sql/query_client_by_id.sql
var queryClientById string

//go:embed sql/insert_clients.sql
var insertClients string

func NewRepository(db *sql.DB) (*repository, error) {
	queryAllClientsStmt, err := db.Prepare(queryAllClients)
	if err != nil {
		return nil, err
	}
	queryClientByIdStmt, err := db.Prepare(queryClientById)
	if err != nil {
		return nil, err
	}

	insertClientsStmt, err := db.Prepare(insertClients)
	if err != nil {
		return nil, err
	}
	return &repository{
		db:              db,
		queryAllClients: queryAllClientsStmt,
		queryClientById: queryClientByIdStmt,
		insertClients:   insertClientsStmt,
	}, nil

}

type repository struct {
	db              *sql.DB
	queryAllClients *sql.Stmt
	queryClientById *sql.Stmt
	insertClients   *sql.Stmt
}
