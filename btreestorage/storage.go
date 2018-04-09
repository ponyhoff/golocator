package btreestorage

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"os"
)

type BTStorage struct {
	MemTree *BTree
}

func New() *BTStorage {
	register()
	s := BTStorage{}
	return &s
}

func FromGOBs(r io.Reader) (*BTStorage, error) {
	register()

	s := BTStorage{}

	dec := gob.NewDecoder(r)

	err := dec.Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (b BTStorage) Query(key string) (interface{}, bool) {
	if b.MemTree == nil {
		return nil, false
	}
	return b.MemTree.find(key)
}

func (b *BTStorage) Put(key string, data interface{}) bool {
	if b.MemTree == nil {
		b.MemTree = &BTree{}
	}

	return b.MemTree.insert(key, data)
}

func (b *BTStorage) WriteTo(w io.Writer) error {
	enc := gob.NewEncoder(w)
	err := enc.Encode(b)
	if err != nil {
		return err
	}

	return nil
}

func (b *BTStorage) DumpToFile(fname string) (int, error) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	// Encode (send) the value.
	err := enc.Encode(b)
	if err != nil {
		log.Fatal("encode error:", err)
		return 0, err
	}

	file, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.WriteString(network.String())

	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
	return bytesWritten, nil
}

func register() {
	gob.Register(BTree{})
	gob.Register(BTStorage{})
}
