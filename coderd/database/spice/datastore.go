package spice

import (
	"context"

	"github.com/authzed/spicedb/pkg/datastore"
	"github.com/authzed/spicedb/pkg/datastore/options"
)

var _ datastore.Datastore = (*DatastoreWrapper)(nil)

type DatastoreWrapper struct {
	// Datastore is the underlying datastore implementation.
	Datastore datastore.Datastore
}

func (d DatastoreWrapper) SnapshotReader(revision datastore.Revision) datastore.Reader {
	return d.Datastore.SnapshotReader(revision)
}

func (d DatastoreWrapper) ReadWriteTx(ctx context.Context, userFunc datastore.TxUserFunc, option ...options.RWTOptionsOption) (datastore.Revision, error) {
	return d.Datastore.ReadWriteTx(ctx, userFunc, option...)
}

func (d DatastoreWrapper) OptimizedRevision(ctx context.Context) (datastore.Revision, error) {
	return d.Datastore.OptimizedRevision(ctx)
}

func (d DatastoreWrapper) HeadRevision(ctx context.Context) (datastore.Revision, error) {
	return d.Datastore.HeadRevision(ctx)
}

func (d DatastoreWrapper) CheckRevision(ctx context.Context, revision datastore.Revision) error {
	return d.Datastore.CheckRevision(ctx, revision)
}

func (d DatastoreWrapper) RevisionFromString(serialized string) (datastore.Revision, error) {
	return d.Datastore.RevisionFromString(serialized)
}

func (d DatastoreWrapper) Watch(ctx context.Context, afterRevision datastore.Revision, options datastore.WatchOptions) (<-chan *datastore.RevisionChanges, <-chan error) {
	return d.Datastore.Watch(ctx, afterRevision, options)
}

func (d DatastoreWrapper) ReadyState(ctx context.Context) (datastore.ReadyState, error) {
	return d.Datastore.ReadyState(ctx)
}

func (d DatastoreWrapper) Features(ctx context.Context) (*datastore.Features, error) {
	return d.Datastore.Features(ctx)
}

func (d DatastoreWrapper) Statistics(ctx context.Context) (datastore.Stats, error) {
	return d.Datastore.Statistics(ctx)
}

func (d DatastoreWrapper) Close() error {
	return d.Datastore.Close()
}
