package main

type Dictionary map[string]string

const (
	ErrorNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists       = DictionaryErr("word already exists in dictionary")
	ErrorWordDoesNotExist = DictionaryErr("word does not exist in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
