package errors_test

import (
	"testing"

	"fmt"

	"github.com/sagikazarmark/utilz/errors"
)

func TestTestHandler_Handle(t *testing.T) {
	handler := errors.NewTestHandler()

	err := fmt.Errorf("internal error")

	handler.Handle(err)

	if got, want := handler.Last(), err; got != want {
		t.Fatalf("expected to log a specific error, received: %v", got)
	}
}

func TestTestHandler_Errors(t *testing.T) {
	handler := errors.NewTestHandler()

	err1 := fmt.Errorf("internal error")
	err2 := fmt.Errorf("internal error")

	handler.Handle(err1)
	handler.Handle(err2)

	if got := handler.Errors(); got[0] != err1 || got[1] != err2 {
		t.Fatalf("expected to log specific errors, received: %v", got)
	}
}

func TestTestHandler_Last(t *testing.T) {
	handler := errors.NewTestHandler()

	err1 := fmt.Errorf("internal error")
	err2 := fmt.Errorf("internal error")

	handler.Handle(err1)
	handler.Handle(err2)

	if got := handler.Last(); got != err2 {
		t.Fatalf("expected to log a specific error, received: %v", got)
	}
}
