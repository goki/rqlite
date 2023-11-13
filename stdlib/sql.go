// Package stdlib is the compatability layer from gorqlite to database/sql.
package stdlib

import (
	"database/sql/driver"

	"github.com/rqlite/gorqlite"
)

type Driver struct{}

func (d *Driver) Open(name string) (driver.Conn, error) {
	conn, err := gorqlite.Open(name)
	if err != nil {
		return nil, err
	}
	return &Conn{conn}, nil
}

type Conn struct {
	*gorqlite.Connection
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *Conn) Close() error {
	c.Connection.Close()
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, nil
}
