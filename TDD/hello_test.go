package main

import "testing"

func TestMain(t *testing.T) {
	got := Hello()
	want := "Hello world!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
