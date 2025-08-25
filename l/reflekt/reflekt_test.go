package reflekt

import (
	"reflect"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

var (
	// String Slice Checker
	ssc = tablewriter.AssertStringSliceGotWant
	// String checker
	sc = tablewriter.AssertStringGotWant
	//Int Checker
	ic = tablewriter.AssertIntGotWant
)

type (
	Person struct {
		Name    string
		Profile Profile
	}

	Profile struct {
		Age  int
		City string
	}
)

func TestWalk(t *testing.T) {
	t.Run("Length Test", func(t *testing.T) {
		expected := "Circle"

		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			ic(t, len(got), 1)
		}
	})

	t.Run("Got expected", func(t *testing.T) {
		expected := "Circle"

		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if got[0] != expected {
			sc(t, got[0], expected)
		}
	})
}

func TestWalkFn(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "Struct with one strong field",
			Input:         struct{ Name string }{Name: "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{Name: "Chris",
				Profile: Profile{Age: 33, City: "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name:    "Chris",
				Profile: Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "maps",
			Input: map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			ExpectedCalls: []string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			ssc(t, got, test.ExpectedCalls)
		})
	}

	t.Run("With Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}

			aChannel <- Profile{34, "Katowice"}

			close(aChannel)
		}()

		var got []string

		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with func", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string

		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
