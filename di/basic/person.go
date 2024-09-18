package main

import (
	"encoding/json"
	"errors"
)

type Person struct {
	Name  string
	Phone string
}

func (p *Person) validate() error {
	if p.Name == "" {
		return errors.New("name missing")
	}

	if p.Phone == "" {
		return errors.New("phone missing")
	}

	return nil
}

func (p *Person) encode() ([]byte, error) {
	return json.Marshal(p)
}
