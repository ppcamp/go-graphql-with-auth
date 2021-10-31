package mocks

import (
	"database/sql"

	"github.com/ppcamp/go-graphql-with-auth/internal/repository/postgres"
	"github.com/stretchr/testify/mock"
)

//#region: Storage

type MockedStorage struct {
	mock.Mock
}

func (m *MockedStorage) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedStorage) SetMaxIdleConns(n int) {
	m.Called(n)
}

func (m *MockedStorage) SetMaxOpenConns(n int) {
	m.Called(n)
}

func (m *MockedStorage) StartTransaction() (postgres.Transaction, error) {
	args := m.Called()

	tr := args.Get(0)
	err := args.Get(1)

	var cerr error
	var ttransaction postgres.Transaction

	if err != nil {
		cerr = err.(error)
	}

	if tr != nil {
		ttransaction = tr.(postgres.Transaction)
	}

	return ttransaction, cerr
}

func (m *MockedStorage) Exec(query string, args ...interface{}) (sql.Result, error) {
	a := m.Called(query, args)

	v := a.Get(0)
	err := a.Get(1)

	var cerr error
	var tv sql.Result

	if err != nil {
		cerr = err.(error)
	}

	if v != nil {
		tv = v.(sql.Result)
	}

	return tv, cerr
}

func (m *MockedStorage) Close() error {
	args := m.Called()

	err := args.Get(0)

	if err != nil {
		return err.(error)
	}

	return nil
}

//#endregion

//#region: Transaction

type MockedTransaction struct {
	mock.Mock
}

func (m *MockedTransaction) Commit() error {
	args := m.Called()

	err := args.Get(0)

	if err != nil {
		return err.(error)
	}
	return nil
}

func (m *MockedTransaction) Rollback() error {
	args := m.Called()

	err := args.Get(0)

	if err != nil {
		return err.(error)
	}

	return nil
}

//#endregion
