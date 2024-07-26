package test

import "testing"

func Eq[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Fatalf("want: %v\ngot: %v", want, got)
	}
}
