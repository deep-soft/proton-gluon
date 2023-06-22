package db_impl

import (
	"github.com/ProtonMail/gluon/db"
	"github.com/ProtonMail/gluon/internal/db_impl/sqlite3"
)

func NewSQLiteDB(options ...sqlite3.Option) db.ClientInterface {
	return sqlite3.NewBuilder(options...)
}
