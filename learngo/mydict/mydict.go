package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound    = errors.New("Not Found")
	errCantUpdate  = errors.New("Cant update non-existing word")
	errWorldExists = errors.New("That word already exists")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exits := d[word]
	if exits {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWorldExists
	}
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return err
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
