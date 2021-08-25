package KV

import (
	"bytes"
	"github.com/dgraph-io/badger/v3"
	"os"
	"testing"
	"time"
)

var ns *NameSpace

func TestNewNameSpace(t *testing.T) {
	ns = NewNameSpace(Config{
		DiskLess: true,
	})

	if _, err := os.Stat(ns.DatabasePath); err == nil || !ns.DiskLess {
		t.Error("diskless is not set correctly.")
	}
}

func TestNameSpace_Put(t *testing.T) {
	if err := ns.Put([]byte("a1"), []byte("b1"), 0); err != nil {
		t.Error("put() returns error:", err)
	}

	if err := ns.Put([]byte("a2"), []byte("b2"), time.Second); err != nil {
		t.Error("put() returns error:", err)
	}

	time.Sleep(time.Second)
}

func TestNameSpace_Get(t *testing.T) {
	b, err := ns.Get([]byte("a1"))
	if err != nil {
		t.Error("get() returns error:", err)
	}

	if bytes.Compare(b, []byte("b1")) != 0 {
		t.Error("get() returns wrong value, expect: b1, get:", string(b))
	}

	b, err = ns.Get([]byte("a2"))
	if err != badger.ErrKeyNotFound {
		t.Error("get() returns wrong error, expect:", badger.ErrKeyNotFound, ", get:", err)
	}
}

func TestNameSpace_Delete(t *testing.T) {
	if err := ns.Delete([]byte("a1")); err != nil {
		t.Error("delete() returns error:", err)
	}

	_, err := ns.Get([]byte("a1"))
	if err != badger.ErrKeyNotFound {
		t.Error("get() returns wrong error, expect:", badger.ErrKeyNotFound, ", get:", err)
	}
}
