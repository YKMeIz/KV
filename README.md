# KV

[![Go Reference](https://img.shields.io/badge/go-reference-informational.svg?style=flat-square)](https://pkg.go.dev/github.com/YKMeIz/KV)
[![Go Report Card](https://goreportcard.com/badge/github.com/YKMeIz/Pill?style=flat-square)](https://goreportcard.com/report/github.com/YKMeIz/KV)
[![License](https://img.shields.io/github/license/YKMeIz/Pill.svg?color=%232b2b2b&style=flat-square)](https://github.com/YKMeIz/KV/blob/master/LICENSE)

KV is a simple key-value store with support of expiration TTL (time to live).


# Get Started

The following example shows how to use KV.

```go

package main

import (
	"fmt"
	"github.com/YKMeIz/KV"
	"time"
)

func main() {
	// Set configuration for key-value store.
	config := KV.Config{
		DiskLess:     false,
		DatabasePath: "",
	}

	// Create a key-value store namespace.
	ns := KV.NewNameSpace(config)

	// Save a key-value with no time limit
	if err := ns.Put([]byte("key1"), []byte("value1"), 0); err != nil {
		panic(err)
	}

	// Save a key-value with time limit
	if err := ns.Put([]byte("key2"), []byte("value2"), 2*time.Second); err != nil {
		panic(err)
	}

	// Read the value of given key - key1
	if b, err := ns.Get([]byte("key1")); err != nil {
		panic(err)
	} else {
		// Will print: value1
		fmt.Println("The value of key1 is", string(b))
	}

	// Read the value of given key - key2
	if b, err := ns.Get([]byte("key2")); err != nil {
		panic(err)
	} else {
		// Will print: value2
		fmt.Println("The value of key2 is", string(b))
	}

	fmt.Println("Wait 2 seconds")
	time.Sleep(2 * time.Second)

	// Try to read a auto-deleted value
	if _, err := ns.Get([]byte("key2")); err != nil {
		fmt.Println("error:", err)
	}

}

```

The terminal will print:

```plain
The value of key1 is value1
The value of key2 is value2
Wait 2 seconds
error: Key not found
```
