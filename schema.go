package KV

import (
	"github.com/dgraph-io/badger/v3"
)

// NameSpace defines a storage space.
type NameSpace struct {
	Config
	db *badger.DB
}

// Config describes configuration of a namespace.
type Config struct {
	// DiskLess is used to set if running in non database filename mode.
	DiskLess bool
	// DatabasePath is used to identify database filename and path.
	// Set unique path to avoid database file collision.
	DatabasePath string
}
