package episode0

import "errors"

var (
	ErrNotFound = errors.New("Not found")
)

type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
