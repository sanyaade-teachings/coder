package spice

import (
	"database/sql"

	"github.com/coder/coder/v2/coderd/database"
)

type SpiceDBTX struct {
	*SpiceDB
}

func (s *SpiceDB) Wrap(tx database.Store) *SpiceDBTX {
	cpy := *s
	cpy.Store = tx
	return &SpiceDBTX{
		SpiceDB: &cpy,
	}
}

func (s *SpiceDB) InTx(f func(database.Store) error, opts *sql.TxOptions) error {
	return s.Store.InTx(func(nestedTX database.Store) error {
		wrapped := s.Wrap(nestedTX)
		return f(wrapped)
	}, opts)
}

func (s *SpiceDBTX) InTx(f func(database.Store) error, opts *sql.TxOptions) error {
	// Do not double wrap transactions.
	return f(s)
}
