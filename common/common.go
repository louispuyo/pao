package main

import (
	"fmt"
)

type User struct {
	ID                  int
	FirstName, LastName string
	Password            string
}

// DATABASE
type Profile struct {
	database map[int]User //private
}

func (c *Profile) Add(payload User, reply *User) error {

	// check if User already exists in the database
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("User with id '%d' already exists", payload.ID)
	}

	// add User `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Get methods returns a User with specific id (procedure).
func (c *Profile) Get(payload int, reply *User) error {

	// get User with id `p` from the database
	result, ok := c.database[payload]

	// check if User exists in the database
	if !ok {
		return fmt.Errorf("User with id '%d' does not exist", payload)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

func (s User) FullName() string {
	return s.FirstName + " " + s.LastName
}

// NewProfile function returns a new instance of Profile (pointer).
func NewProfile() *Profile {
	return &Profile{
		database: make(map[int]User),
	}
}
