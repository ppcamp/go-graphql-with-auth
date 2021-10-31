package postgres

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	cfg "github.com/ppcamp/go-graphql-with-auth/internal/config"
)

var (
	ErrEmptyResult   = sql.ErrNoRows
	ErrDuplicatedRow = errors.New("duplicated row")
)

//#region: Store
type store struct {
	*sqlx.DB
}

// A store is the base class used to connection
func newStore() (str *store, err error) {
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.Url)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(cfg.Database.ConnIdleMaxTime)

	str = &store{db}
	str.SetMaxIdleConns(cfg.Database.MinConnections)
	str.SetMaxOpenConns(cfg.Database.MaxConnections)
	return
}

//#endregion

//#region: storage
type storage struct {
	*store
}

// Create a new transaction to the current storage
func (s *storage) StartTransaction() (Transaction, error) {
	tx, err := s.Beginx()

	if err != nil {
		return nil, err
	}

	return NewTransaction(tx), nil
}

// Setup a new postgres database connection
func NewStorage() (Storage, error) {
	store, err := newStore()
	if err != nil {
		return nil, err
	}

	return &storage{store: store}, nil
}

//#endregion

// Storage is just a wrapper that we used to give us more control when
// doing mocks
type Storage interface {
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Close() error

	Exec(query string, args ...interface{}) (sql.Result, error)

	StartTransaction() (Transaction, error)
}
