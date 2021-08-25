package KV

import "github.com/dgraph-io/badger/v3"

// Delete removes given key from namespace.
func (ns *NameSpace) Delete(key []byte) error {
	return ns.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}
