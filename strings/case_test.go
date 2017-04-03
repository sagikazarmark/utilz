package strings_test

import (
	"testing"

	"github.com/sagikazarmark/utilz/strings"
)

func TestToSnake(t *testing.T) {
	data := map[string]string{
		"foo":       "foo",
		"FooBar":    "foo_bar",
		"fooBar":    "foo_bar",
		"Foo-Bar":   "foo_bar",
		"Foo Bar":   "foo_bar",
		"FOOBar":    "foo_bar",
		"FOOBarBaz": "foo_bar_baz",
		"FOOBarBAZ": "foo_bar_baz",
		"Foo_-Bar":  "foo__bar",
	}

	for in, want := range data {
		if got := strings.ToSnake(in); got != want {
			t.Errorf("converting '%s' to snake case failed, expected: %s, actual: %s", in, want, got)
		}
	}
}
