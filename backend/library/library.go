package library

import (
	"database/sql"
)

type Store struct {
	RawDB *sql.DB
}
