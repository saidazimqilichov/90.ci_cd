package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Test"
	want := "Hello Test"

	if got := `Hi ${name}`; got != want {
		t.Errorf("sayHello(%q) = %q, want %q", name, got, want)
	}
}