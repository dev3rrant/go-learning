package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("word does not exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("blah")
		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "another"
		definition := "test definition"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, "another", "test definition")
	})
	t.Run("existing word", func(t *testing.T) {
		word := "another"
		definition := "test definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, "another", "test definition")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "another"
		definition := "test definition"
		updatedDefinition := "update definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, updatedDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updatedDefinition)
	})
	t.Run("update new word", func(t *testing.T) {
		word := "another"
		differentWord := "test"
		definition := "test definition"
		updatedDefinition := "update definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(differentWord, updatedDefinition)

		assertError(t, err, ErrorWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrorNotFound {
			t.Errorf("expected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
