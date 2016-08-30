package tests

import "fmt"

func convert(i int) (string, error) {
	switch i {
	case 1:
		return "one", nil
	case 2, 3:
		return "two or three", nil
	default:
		return "", fmt.Errorf("unable to handle: %v", i)
	}
}

// ConvertToString converts the given integer (i) to a string
// i can not be greater than 3
func ConvertToString(i int) (string, error) {
	if i > 3 {
		return "", fmt.Errorf("i can not be greater then 3: got %v", i)
	}
	return convert(i)
}
