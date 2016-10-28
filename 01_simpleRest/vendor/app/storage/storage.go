package storage

import "errors"

type Database struct {
	entries map[string]string
}

func NewInMemoryDatabase() Database {
	return Database{entries: make(map[string]string)}
}

func (db Database) Set(key string, value string) {
	db.entries[key] = value
}

func (db Database) Get(key string) (string, error) {
	val, ok := db.entries[key]
	if !ok {
		return "", errors.New("nothing found")
	}
	return val, nil
}
