package strings_test

import (
	"testing"

	"github.com/sagikazarmark/utilz/strings"
)

func TestToCamel(t *testing.T) {
	data := map[string]string{
		"foo": "Foo",

		// Snake
		"_foo":        "_Foo",
		"__foo":       "_Foo",
		"___foo":      "__Foo",
		"foo_":        "Foo_",
		"foo__":       "Foo_",
		"foo___":      "Foo__", // TODO: good idea?
		"foo_bar":     "FooBar",
		"foo_bar_baz": "FooBarBaz",
		"foo__bar":    "Foo_Bar",

		// Spinal
		"-foo":        "-Foo",
		"--foo":       "-Foo",
		"foo-":        "Foo-",
		"foo--":       "Foo-",
		"foo---":      "Foo--", // TODO: good idea?
		"foo-bar":     "FooBar",
		"foo-bar-baz": "FooBarBaz",
		"foo--bar":    "Foo-Bar",

		// Train
		"-Foo":        "-Foo",
		"--Foo":       "-Foo",
		"---Foo":      "--Foo",
		"Foo-":        "Foo-",
		"Foo--":       "Foo-",
		"Foo---":      "Foo--", // TODO: good idea?
		"Foo-Bar":     "FooBar",
		"Foo-Bar-Baz": "FooBarBaz",
		"Foo--Bar":    "Foo-Bar",
	}

	for in, want := range data {
		if got := strings.ToCamel(in); got != want {
			t.Errorf("converting '%s' to camel case failed, expected: %s, actual: %s", in, want, got)
		}
	}
}

func TestToSnake(t *testing.T) {
	data := map[string]string{
		"foo":       "foo",
		"FooBar":    "foo_bar",
		"fooBar":    "foo_bar",
		"Foo_Bar":   "foo_bar",
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

func TestToSpinal(t *testing.T) {
	data := map[string]string{
		"foo":       "foo",
		"FooBar":    "foo-bar",
		"fooBar":    "foo-bar",
		"Foo_Bar":   "foo-bar",
		"Foo-Bar":   "foo-bar",
		"Foo Bar":   "foo-bar",
		"FOOBar":    "foo-bar",
		"FOOBarBaz": "foo-bar-baz",
		"FOOBarBAZ": "foo-bar-baz",
		"Foo_-Bar":  "foo--bar",
	}

	for in, want := range data {
		if got := strings.ToSpinal(in); got != want {
			t.Errorf("converting '%s' to spinal case failed, expected: %s, actual: %s", in, want, got)
		}
	}
}

func TestToTrain(t *testing.T) {
	data := map[string]string{
		"foo":       "Foo",
		"FooBar":    "Foo-Bar",
		"fooBar":    "Foo-Bar",
		"Foo_Bar":   "Foo-Bar",
		"Foo-Bar":   "Foo-Bar",
		"Foo Bar":   "Foo-Bar",
		"FOOBar":    "Foo-Bar",
		"FOOBarBaz": "Foo-Bar-Baz",
		"FOOBarBAZ": "Foo-Bar-Baz",
		"Foo_-Bar":  "Foo--Bar",
	}

	for in, want := range data {
		if got := strings.ToTrain(in); got != want {
			t.Errorf("converting '%s' to train case failed, expected: %s, actual: %s", in, want, got)
		}
	}
}
