package aps

import (
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}
		errDelete := dictionary.Delete(word)
		tablewriter.AssertErrorType(t, errDelete, nil)
	})

	t.Run("non-existing-word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)
		tablewriter.AssertErrorType(t, err, ErrWordNotExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		tablewriter.AssertErrorType(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}
		err := dict.Update(word, definition)
		tablewriter.AssertErrorType(t, err, ErrWordNotExists)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "definition of test"}

	t.Run("Known Word", func(t *testing.T) {
		got, errSearch := dictionary.Search("test")
		want := "definition of test"

		tablewriter.AssertStringGotWant(t, got, want)
		tablewriter.AssertErrorNil(t, errSearch)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := "could not find the word you are looking for"

		tablewriter.AssertStringGotWant(t, got, "")

		if err == nil {
			t.Fatalf("Expected to get an error")
		}

		tablewriter.AssertStringGotWant(t, err.Error(), want)
	})

	t.Run("Unknown Word Refector v1", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		tablewriter.AssertErrorType(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		want := "this is just a test"

		tablewriter.AssertErrorType(t, err, nil)
		assertDefinition(t, dictionary, "test", want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: word, definition: definition}
		err := dictionary.Add(word, "new test")
		tablewriter.AssertErrorType(t, err, ErrWordExists)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	tablewriter.AssertStructGotWant(t, got, definition)
}
