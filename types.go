package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// User represents a user of this software
type User struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	First   string `json:"first"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
	Pass    []byte
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID      int       `json:"id"`
	User    int       `json:"user_id"`
	Amount  float32   `json:"amount"`
	Tags    []string  `json:"tags"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
}

// Save saves this Transaction to the database
func (t *Transaction) Save(db *DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		id, _ := b.NextSequence()
		t.ID = int(id)
		t.Created = time.Now()
		t.Updated = time.Now()
		tj, err := json.Marshal(t)
		if err != nil {
			return errors.New("unable to serialize transaction to json")
		}
		err = b.Put(itob(t.ID), tj)
		if err != nil {
			return errors.New("unable to save this transaction")
		}
		return nil
	})
}

// Delete removes this Transaction from the database
func (t *Transaction) Delete(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		err := b.Delete(itob(t.ID))
		if err != nil {
			return errors.New("unable to delete this transaction")
		}
		return nil
	})
}