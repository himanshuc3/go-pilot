package dictionary

type Dictionary map[string]string

// NOTE:
// 1. Mr Dave Chaney's error handling guide is a must read:
// https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
// 2. Apparently the golden standard is not to return sentinel errors and create dependencies between packages.
// var ErrNotFound = errors.New("404: word doesn't exist in the dictionary")
// var ErrWordExists = errors.New("word already exists in the dictionary")
// var ErrDictionaryIsNil = errors.New("dictionary is nil")

// NOTE:
// 1. Constants are like enums in other languages.
// 2. Error constants require implementing a custom error type.
// 3. Mr Dave Chaney has written it all:
// https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrNotFound         = DictionaryErr("404: word doesn't exist in the dictionary")
	ErrWordExists       = DictionaryErr("word already exists in the dictionary")
	ErrDictionaryIsNil  = DictionaryErr("dictionary is nil")
	ErrWordDoesNotExist = DictionaryErr("No-Op: word does not exist in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// NOTE:
// 1. Maps in golang are again passed by value like slices
// but the value being copied is the pointer to the map.
// 2. Gotcha: nil maps (will act as empty maps when reading and cause a runtime panic
// when writing).
// 3. 🚫 var m map[string]string -> creates a nil map
// 4. ✅ var dictionary = make(map[string]string) or var dictionary = map[string]string{}
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
