package testing

import (
	"reflect"
	"slices"
	"testing"
)

// Checks are assertions that can be used in testing to make the tests more readable.
func TestWithChecks(t *testing.T) {
	// For each test we can define the check type internally so we don't collide with any other testing files in
	// the same package. Each check signature may be different based on the values accepted as parameters which
	// are directly related to the functionality being tested.
	type check func(*testing.T, []int)

	// This is a helper function that prevents us from having to instantiate an array in our table tests
	checks := func(cs ...check) []check { return cs }

	// Define some checks
	isEqual := func(expected []int) check {
		return func(t *testing.T, actual []int) {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("expected: %v, actual %v", expected, actual)
			}
		}
	}
	hasLength := func(expected int) check {
		return func(t *testing.T, array []int) {
			if len(array) != expected {
				t.Errorf("result has length %d, expected %d", len(array), expected)
			}
		}
	}

	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		checks []check
	}{
		{
			name: "check sorting",
			args: args{
				nums: []int{5, 2, 4, 3, 1},
			},
			checks: checks(
				isEqual([]int{1, 2, 3, 4, 5}),
				hasLength(5),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slices.Sort(tt.args.nums)
			for _, ch := range tt.checks {
				ch(t, tt.args.nums)
			}
		})
	}
}
