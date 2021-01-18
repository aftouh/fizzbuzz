package core

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func fizzbuzz16() []string {
	return []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}
}
func TestFizzbuzz(t *testing.T) {
	tt := []struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
		err   error
		res   []string
	}{
		{"base version 16", 3, 5, 16, "fizz", "buzz", nil, fizzbuzz16()},
		{"error limit", 3, 5, 0, "", "", errors.New("limit must be great or equal to 1"), []string{}},
		{"error int1 0", 0, 1, 1, "", "", errors.New("int1 and int2 must be greater than 0"), []string{}},
		{"error int2 0", 1, 0, 0, "", "", errors.New("int1 and int2 must be greater than 0"), []string{}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Fizzbuzz(tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)
			if tc.err != nil { // Check exptected error
				if err == nil {
					t.Errorf("Missed an expected error: %v", tc.err)
				}
				if err.Error() != tc.err.Error() {
					t.Errorf("The obtained err %q doesn't match the exptected one %q", err, tc.err)
				}
				return
			}

			// Check non error result
			if err != nil {
				t.Errorf("Get an unexpected error: %v", tc.err)
			}

			if !cmp.Equal(res, tc.res) {
				t.Errorf("The get result %q doesn't match the exptected one %q", res, tc.res)
			}
		})
	}
}
