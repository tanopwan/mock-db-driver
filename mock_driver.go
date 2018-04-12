package test

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

type mockDriver struct {
	sync.Mutex
}

// Register ... for pgmock driver
func Register() {
	sql.Register("pgmock", &mockDriver{})
}

func init() {
	Register()
}

func (d *mockDriver) Open(dsn string) (driver.Conn, error) {
	d.Lock()
	defer d.Unlock()

	c := mockConn{}
	return c, nil
}
