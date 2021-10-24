package maps

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("key not found")
var ErrKeyAlreadyExists = errors.New("key already exists")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	if _, err := d.Search(key); err == nil {
		return ErrKeyAlreadyExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if _, err := d.Search(key); err != nil {
		return ErrKeyNotFound
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	if _, err := d.Search(key); err != nil {
		return ErrKeyNotFound
	}
	delete(d, key)
	return nil
}
