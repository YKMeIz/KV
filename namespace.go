package KV

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

// NewNameSpace creates a new namespace for storing key-value data.
func NewNameSpace(c Config) *NameSpace {
	var (
		ns  NameSpace
		err error
	)

	if c.DiskLess {
		c.DatabasePath = ""
	} else {
		if c.DatabasePath == "" {
			c.DatabasePath = "/tmp/KV.db"
		}
	}

	ns.db, err = badger.Open(badger.DefaultOptions(c.DatabasePath).WithInMemory(c.DiskLess).WithLoggingLevel(badger.ERROR))
	if err != nil {
		log.Fatal(err)
	}

	ns.Config = c

	return &ns
}
