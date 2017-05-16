package util_test

import (
	"testing"

	"fmt"

	"github.com/sagikazarmark/utilz/errors"
	"github.com/sagikazarmark/utilz/util"
)

func CreateShutdownFunc() (util.ShutdownHandler, *bool) {
	var called bool

	f := util.ShutdownFunc(func() {
		called = true
	})

	return f, &called
}

func CreateOrderedShutdownFuncs(num int) ([]util.ShutdownHandler, *[]int) {
	funcs := make([]util.ShutdownHandler, num)
	called := []int{}

	for index := 0; index < num; index++ {
		funcs[index] = func(index int) util.ShutdownHandler {
			return util.ShutdownFunc(func() {
				called = append(called, index+1)
			})
		}(index)
	}

	return funcs, &called
}

func TestShutdownFunc_CallsUnderlyingFunc(t *testing.T) {
	f, called := CreateShutdownFunc()

	var err error

	if got, want := f(), err; got != want {
		t.Fatalf("wrapped functions are expected to return nil, error received: %v", got)
	}

	if *called != true {
		t.Fatal("the wrapped function is expected to be called")
	}
}

func TestShutdownFunc_RecoversErrorPanic(t *testing.T) {
	err := fmt.Errorf("internal error")

	f := util.ShutdownFunc(func() {
		panic(err)
	})

	if got, want := f(), err; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}

func TestShutdownFunc_RecoversStringPanic(t *testing.T) {
	f := util.ShutdownFunc(func() {
		panic("internal error")
	})

	if got, want := f().Error(), "internal error"; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}

func TestShutdownFunc_RecoversAnyPanic(t *testing.T) {
	f := util.ShutdownFunc(func() {
		panic(123)
	})

	if got, want := f().Error(), "Unknown panic, received: 123"; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}

func TestNewShutdownManager(t *testing.T) {
	shutdownManager := util.NewShutdownManager(nil)

	// Test falling back to NullHandler
	shutdownManager.Register(func() error {
		return fmt.Errorf("error")
	})

	shutdownManager.Shutdown()
}

func TestShutdownManager_Register(t *testing.T) {
	f, called := CreateShutdownFunc()

	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	shutdownManager.Register(f)
	shutdownManager.Shutdown()

	if *called != true {
		t.Fatal("the shutdown handler is expected to be called")
	}
}

func TestShutdownManager_Register_ExecutedInOrder(t *testing.T) {
	funcs, called := CreateOrderedShutdownFuncs(2)

	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	shutdownManager.Register(funcs[0], funcs[1])
	shutdownManager.Shutdown()

	if got, want := (*called)[0], 1; got != want {
		t.Fatal("the first shutdown handler is expected to be called first")
	}

	if got, want := (*called)[1], 2; got != want {
		t.Fatal("the second shutdown handler is expected to be called second")
	}
}

func TestShutdownManager_RegisterAsFirst(t *testing.T) {
	funcs, called := CreateOrderedShutdownFuncs(2)

	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	shutdownManager.Register(funcs[1])
	shutdownManager.RegisterAsFirst(funcs[0])
	shutdownManager.Shutdown()

	if got, want := (*called)[0], 1; got != want {
		t.Fatal("the first shutdown handler is expected to be called first")
	}

	if got, want := (*called)[1], 2; got != want {
		t.Fatal("the second shutdown handler is expected to be called second")
	}
}

func TestShutdownManager_Shutdown(t *testing.T) {
	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	shutdownManager.Shutdown()

	if handler.Last() != nil {
		t.Fatal("shutting down not emit an error")
	}
}

func TestShutdownManager_Shutdown_HandleErrors(t *testing.T) {
	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	err := fmt.Errorf("error")

	shutdownManager.Register(func() error {
		return err
	})

	shutdownManager.Shutdown()

	if handler.Last() != err {
		t.Fatal("errors ocurred during shutdown should be handled")
	}
}

func TestShutdownManager_Shutdown_RecoverFromPanic(t *testing.T) {
	handler := errors.NewTestHandler()
	shutdownManager := util.NewShutdownManager(handler)

	err := fmt.Errorf("error")

	func() {
		defer shutdownManager.Shutdown()

		func() {
			panic(err)
		}()
	}()

	if handler.Last() != err {
		t.Fatal("errors ocurred during shutdown should be handled")
	}
}
