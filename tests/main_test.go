package tests

import "testing"

func TestConvert(t *testing.T) {
	in := 1
	out := "one"
	got, err := convert(in)

	if err != nil {
		t.Errorf("Error: %v, Expected: nil", err)
	}

	if got != out {
		t.Errorf("convert(%v): Got: %v, Expected: %v", in, got, out)
	}
}
