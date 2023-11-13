package sqlite

import (
	"context"
	"database/sql"

	"github.com/rqlite/gorqlite"
)

// ConnPool wraps [gorqlite.Connection] to implement [gorm.ConnPool]
type ConnPool struct {
	*gorqlite.Connection
}

func (cp *ConnPool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil
}

func (cp *ConnPool) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return nil, nil
}

func (cp *ConnPool) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	stmt := gorqlite.ParameterizedStatement{Query: query, Arguments: args}
	qr, err := cp.Connection.QueryOneParameterizedContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	return qr, nil
}

func (cp *ConnPool) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	stmt := gorqlite.ParameterizedStatement{Query: query, Arguments: args}
	qr, err := cp.Connection.QueryOneParameterizedContext(ctx, stmt)
	return qr
}
