package underscore

import (
	"fmt"
	"testing"
)

func TestCamel(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithA", "ends_with_a"},
	}

	for _, tc := range tests {
		t.Run(tc.arg, func(t *testing.T) {
			if got := Camel(tc.arg); got != tc.want {
				t.Fatalf("Camel() = %v, want %v", got, tc.want)
			}
			fmt.Println("this won't print upon failure")
		})
	}
}
