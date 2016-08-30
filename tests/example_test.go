package tests_test

import (
	"fmt"
	"testing"

	"github.com/otm/intro/tests"
)

func TestConvertToString(t *testing.T) {
	i := 2
	str, err := tests.ConvertToString(i)
	if err != nil {
		t.Errorf("Got error %v, expected nil", err)
	}

	if str != "two or three" {
		t.Errorf("Got %v, Expected: %v", i, str)
	}
}

func ExampleConvertToString() {
	str, _ := tests.ConvertToString(2)
	fmt.Println(str)
	// Output:
	// two or three
}
