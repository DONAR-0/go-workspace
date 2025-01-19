package aps

import (
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		tablewriter.AssertErrorNil(t, err)
		tablewriter.AssertStringGotWant(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("not a word")
		want := ErrNotFound.Error()
		tablewriter.AssertError(t, err, want)
		tablewriter.AssertStringGotWant(t, got, "")
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("add word", func(t *testing.T) {
		errAdd := dictionary.Add("some word", "some definition")
		tablewriter.AssertErrorNil(t, errAdd)

		got, errSearch := dictionary.Search("some word")
		want := "some definition"

		tablewriter.AssertErrorNil(t, errSearch)
		tablewriter.AssertStringGotWant(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is sparta"}

		newDefinition := "kicking a leg"

		errUpdate := dictionary.Update("test", newDefinition)
		tablewriter.AssertErrorNil(t, errUpdate)

		got, err := dictionary.Search("test")
		tablewriter.AssertErrorNil(t, err)
		tablewriter.AssertStringGotWant(t, got, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		newDefinition := "kicking a leg"

		errUpdate := dictionary.Update("test", newDefinition)
		tablewriter.AssertErrorNil(t, errUpdate)

		got, errSearch := dictionary.Search("test")
		tablewriter.AssertErrorNil(t, errSearch)
		tablewriter.AssertStringGotWant(t, got, newDefinition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dic := Dictionary{word: "test definition"}
		err := dic.Delete(word)
		tablewriter.AssertErrorNil(t, err)
	})
}
