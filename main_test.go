package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting()
	want := "Hello from go-app!"
	if got != want {
		t.Errorf("greeting() = %q, want %q", got, want)
	}
}
