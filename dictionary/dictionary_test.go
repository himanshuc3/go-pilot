package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// NOTE:
	// 1. keys in maps should be comparable for get and set operations to be satisfied.
	dictionary := Dictionary{
		"test": "this is a sample",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is a sample"
		assertStrings(t, got, expected)

	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("test101")
		assertError(t, err, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {

	t.Run("add word", func(t *testing.T) {

		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		expected := "this is just a test"
		assertDefinition(t, dictionary, "test", expected)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		// NOTE:
		// 1. We do have dynamic key value pairs in maps in golang (that's a relief of the chest)
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertNoError(t, err)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one: %s", got)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()
	if got != expected {
		t.Errorf("got error %q expected %q", got, expected)
	}
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertNoError(t, err)
	assertStrings(t, got, definition)
}
