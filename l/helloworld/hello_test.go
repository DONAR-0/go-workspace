package helloworld

import (
	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		name, language string
	}

	tests := []struct {
		name string
		got  args
		want string
	}{
		{"v1", args{name: "World", language: "English"}, "Hello, World"},
		{"v2", args{name: "Chris", language: "English"}, "Hello, Chris"},
		{"v3", args{name: "Elodie", language: "Spanish"}, "Hola, Elodie"},
		{"v4", args{name: "Elodie", language: "French"}, "Bonjour, Elodie"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Hello(test.got.name, test.got.language)
			tablewriter.AssertStringGotWant(t, got, test.want)
		})
	}

	t.Run("v2", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		tablewriter.AssertStringGotWant(t, got, want)
	})

	t.Run("v3", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		tablewriter.AssertStringGotWant(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		tablewriter.AssertStringGotWant(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		tablewriter.AssertStringGotWant(t, got, want)
	})
}
