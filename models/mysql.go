package models

import "database/sql"

// DbQuery ...
type DbConnInteface interface {
	GetResult(*sql.Rows, *string) error
}
