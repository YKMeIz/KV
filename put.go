package KV

import (
	"github.com/dgraph-io/badger/v3"
	"time"
)

// Put stores given key and value in the namespace.
// Stored data will be deleted after ttl period.
// Set ttl to 0 for disabling auto-delete.
func (ns *NameSpace) Put(key, val []byte, ttl time.Duration) error {
	return ns.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, val)
		if ttl.Seconds() != 0 {
			e.WithTTL(ttl)
		}
		return txn.SetEntry(e)
	})
}
