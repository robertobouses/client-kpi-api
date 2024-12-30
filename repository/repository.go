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

//go:embed sql/update_client_by_id.sql
var updateClientById string

//go:embed sql/delete_clients_by_id.sql
var deleteClientsById string

//go:embed sql/client_exists_by_id.sql
var clientExistsById string

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
	updateClientByIdStmt, err := db.Prepare(updateClientById)
	if err != nil {
		return nil, err
	}
	deleteClientsByIdStmt, err := db.Prepare(deleteClientsById)
	if err != nil {
		return nil, err
	}
	clientExistsByIdStmt, err := db.Prepare(clientExistsById)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:                db,
		queryAllClients:   queryAllClientsStmt,
		queryClientById:   queryClientByIdStmt,
		insertClients:     insertClientsStmt,
		updateClientById:  updateClientByIdStmt,
		deleteClientsById: deleteClientsByIdStmt,
		clientExistsById:  clientExistsByIdStmt,
	}, nil

}

type repository struct {
	db                *sql.DB
	queryAllClients   *sql.Stmt
	queryClientById   *sql.Stmt
	insertClients     *sql.Stmt
	updateClientById  *sql.Stmt
	deleteClientsById *sql.Stmt
	clientExistsById  *sql.Stmt
}
