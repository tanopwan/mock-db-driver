package test

import "database/sql/driver"

type mockConn struct {
}

func (c mockConn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c mockConn) Close() error {
	return nil
}

func (c mockConn) Begin() (driver.Tx, error) {
	return nil, nil
}
