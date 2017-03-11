package str_test

import (
	"testing"

	"github.com/sagikazarmark/utilz/str"
)

func TestToSnake(t *testing.T) {
	data := map[string]string{
		"FooBar":  "foo_bar",
		"fooBar":  "foo_bar",
		"Foo-Bar": "foo_bar",
		"Foo Bar": "foo_bar",
		"FOOBar":  "foo_bar",
	}

	for in, want := range data {
		got := str.ToSnake(in)

		if got != want {
			t.Errorf("converting '%s' to snake case failed, expected: %s, actual: %s", in, want, got)
		}
	}
}
