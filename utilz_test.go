package utilz_test

import (
	"errors"
	"testing"

	"github.com/sagikazarmark/utilz"
)

func TestMust_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	utilz.Must(errors.New("should panic"))
}

func TestMust_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not panic")
		}
	}()

	utilz.Must(nil)
}
