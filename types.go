package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// User represents a user of this software
type User struct {
	ID      int       `json:"id" gorm:"primary_key"`
	Login   string    `json:"login"`
	First   string    `json:"first"`
	Pass    []byte    `json:"pass"`
	Created time.Time `json:"created" gorm:"DEFAULT:current_timestamp"`
	Updated time.Time `json:"updated" gorm:"DEFAULT:current_timestamp"`
	Deleted time.Time `json:"deleted"`
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID      int       `json:"id" gorm:"primary_key"`
	User    int       `json:"user_id"`
	Amount  float32   `json:"amount"`
	Tags    []Tag     `json:"tags" gorm:"many2many:transactions_tags;"`
	Note    string    `json:"note"`
	Created time.Time `json:"created" gorm:"DEFAULT:current_timestamp"`
	Updated time.Time `json:"updated" gorm:"DEFAULT:current_timestamp"`
	Deleted time.Time `json:"deleted"`
}

// Tag is basically just a string
type Tag struct {
	ID      int       `json:"id" gorm:"primary_key"`
	Name    string    `json:"name"`
	Created time.Time `json:"created" gorm:"DEFAULT:current_timestamp"`
	Updated time.Time `json:"updated" gorm:"DEFAULT:current_timestamp"`
	Deleted time.Time `json:"deleted"`
}

// Create saves this Transaction to the database
func (t *Transaction) Create(db *gorm.DB) {
	db.Create(t)
}

// Command line options packed together
type params struct {
	port   int
	dbhost string
	dbport int
	dbuser string
	dbname string
	dbpass string
}

func (p *params) validate() error {
	if p.dbpass == "foo" {
		// TODO: CLEANUP
		return nil
		return errors.New("database password not provided")
	}
	return nil
}
