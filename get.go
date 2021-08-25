package KV

import "github.com/dgraph-io/badger/v3"

// Get obtains value with given key from the namespace.
func (ns *NameSpace) Get(key []byte) ([]byte, error) {
	var val []byte
	if err := ns.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	}); err != nil {
		return nil, err
	}
	return val, nil
}
