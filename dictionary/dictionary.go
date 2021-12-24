package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

//thin wrapper around map
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) {
	//An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
	//https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
	//A map value is a pointer to a runtime.hmap structure.

	d[word] = definition

}
