package auth

import (
	"github.com/eturella/go-mysql-test/sql"

	"vitess.io/vitess/go/mysql"
)

// None is an Auth method that always succeeds.
type None struct{}

// Mysql implements Auth interface.
func (n *None) Mysql() mysql.AuthServer {
	return new(mysql.AuthServerNone)
}

// Allowed ??????
func (n *None) Allowed(ctx *sql.Context, permission Permission) error {
	return nil
}
