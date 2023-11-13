// Package stdlib is the compatability layer from gorqlite to database/sql.
package stdlib

import (
	"database/sql"
	"database/sql/driver"

	"github.com/rqlite/gorqlite"
)

func init() {
	sql.Register("rqlite", &Driver{})
}

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

type Stmt struct {
	Stmt string
	Conn *Conn
}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	return -1
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	wr, err := s.Conn.WriteOne(s.Stmt)
	if err != nil {
		return nil, err
	}
	return &Result{wr}, nil
}

type Result struct {
	gorqlite.WriteResult
}

func (r *Result) LastInsertId() (int64, error) {
	return r.WriteResult.LastInsertID, r.WriteResult.Err
}

func (r *Result) RowsAffected() (int64, error) {
	return r.WriteResult.RowsAffected, r.WriteResult.Err
}
