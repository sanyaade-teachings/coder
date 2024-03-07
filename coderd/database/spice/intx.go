package spice

import (
	"database/sql"

	"github.com/coder/coder/v2/coderd/database"
)

type reverter interface {
	AddRevert(func())
	RevertAll()
}

type SpiceDBTX struct {
	*SpiceDB
	Reverts []func()
}

func (s *SpiceDB) Wrap(tx database.Store) *SpiceDBTX {
	cpy := *s
	cpy.Store = tx
	spiceTx := &SpiceDBTX{
		SpiceDB: &cpy,
	}
	spiceTx.SpiceDB.reverts = spiceTx
	return spiceTx
}

func (s *SpiceDB) InTx(f func(database.Store) error, opts *sql.TxOptions) error {
	return s.Store.InTx(func(nestedTX database.Store) error {
		wrapped := s.Wrap(nestedTX)
		err := f(wrapped)
		if err != nil {
			s.reverts.RevertAll()
			return err
		}
		return nil
	}, opts)
}

func (s *SpiceDBTX) AddRevert(f func()) {
	s.Reverts = append(s.Reverts, f)
}

func (s *SpiceDBTX) RevertAll() {
	all := s.Reverts
	s.Reverts = nil

	for i := range all {
		all[i]()
	}
}

func (s *SpiceDBTX) InTx(f func(database.Store) error, _ *sql.TxOptions) error {
	// Do not double wrap transactions.
	return f(s)
}

func noop() {}
