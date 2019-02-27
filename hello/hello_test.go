package hello

import "testing"

func TestFarewell(t *testing.T) {
	want := "goodbye";
	if Sup() != want {
		t.Errorf("Goodbye() != %q", want);
	}
}