package maps

import "errors"

// ErrNotFound is thrown when the given key can't be looked up in the dictionary
var ErrNotFound = errors.New("Could not find the word you were looking for")

// Dictionary wraps the built-in map[string]string and adds some utility functions
type Dictionary map[string]string

// Search searches for the value stored in the given dictionary under the given key
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

// Add adds the given key and value to the dictionary
func (d Dictionary) Add(key string, value string) {
	d[key] = value
}

// Delete deletes a key and underlying value from the dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
