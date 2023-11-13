// Package stdlib is the compatability layer from gorqlite to database/sql.
package stdlib

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log/slog"

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
	a := make([]any, len(args))
	for i, v := range args {
		a[i] = v
	}
	stmt := gorqlite.ParameterizedStatement{Query: s.Stmt, Arguments: a}
	wr, err := s.Conn.WriteOneParameterized(stmt)
	if err != nil {
		return &Result{wr}, err
	}
	return &Result{wr}, nil
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	a := make([]any, len(args))
	for _, v := range args {
		if v.Name != "" {
			slog.Error("rqlite: Stmt.ExecContext: rqlite sql driver does not support named parameters, but got one", "name", v.Name, "value", v.Value)
		}
		a[v.Ordinal] = v.Value
	}
	stmt := gorqlite.ParameterizedStatement{Query: s.Stmt, Arguments: a}
	wr, err := s.Conn.WriteOneParameterizedContext(ctx, stmt)
	if err != nil {
		return &Result{wr}, err
	}
	return &Result{wr}, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	a := make([]any, len(args))
	for i, v := range args {
		a[i] = v
	}
	stmt := gorqlite.ParameterizedStatement{Query: s.Stmt, Arguments: a}
	qr, err := s.Conn.QueryOneParameterized(stmt)
	if err != nil {
		return &Rows{qr}, err
	}
	return &Rows{qr}, nil
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	a := make([]any, len(args))
	for _, v := range args {
		if v.Name != "" {
			slog.Error("rqlite: Stmt.QueryContext: rqlite sql driver does not support named parameters, but got one", "name", v.Name, "value", v.Value)
		}
		a[v.Ordinal] = v.Value
	}
	stmt := gorqlite.ParameterizedStatement{Query: s.Stmt, Arguments: a}
	qr, err := s.Conn.QueryOneParameterizedContext(ctx, stmt)
	if err != nil {
		return &Rows{qr}, err
	}
	return &Rows{qr}, nil
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

type Rows struct {
	gorqlite.QueryResult
}

func (r *Rows) Columns() []string {
	return r.QueryResult.Columns()
}

func (r *Rows) Close() error {
	return r.Err
}

func (r *Rows) Next(dest []driver.Value) error {
	ok := r.QueryResult.Next()
	if !ok {
		return nil
	}
	a := make([]any, len(dest))
	for i, v := range dest {
		a[i] = v
	}
	return r.QueryResult.Scan(a...)
}
