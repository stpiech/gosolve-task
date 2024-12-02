package search

import (
	"reflect"
	"testing"
)

func TestFindIndexOrClosest(t *testing.T) {
	type input struct {
		numbers []int
		target  int
	}

	tests := []struct {
		name string
		in   input
		out  SearchResult
		err  bool
	}{
		{
			name: "When empty list is passed",
			in:   input{numbers: []int{}, target: 3},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When list has one element which is a target",
			in:   input{numbers: []int{3}, target: 3},
			out:  SearchResult{Index: 0, Number: 3},
			err:  false,
		},
		{
			name: "When list has one element which is close to target",
			in:   input{numbers: []int{32}, target: 33},
			out:  SearchResult{Index: 0, Number: 32},
			err:  false,
		},
		{
			name: "When list has one element which is no close to target",
			in:   input{numbers: []int{5}, target: 33},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When closest lower element in 29.7 and number in array is 30",
			in:   input{numbers: []int{30}, target: 33},
			out:  SearchResult{Index: 0, Number: 30},
			err:  false,
		},
		{
			name: "When closest lower element in 29.7 and number in array is 29",
			in:   input{numbers: []int{29}, target: 33},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When closest higher element in 36.3 and number in array is 36",
			in:   input{numbers: []int{36}, target: 33},
			out:  SearchResult{Index: 0, Number: 36},
			err:  false,
		},
		{
			name: "When closest higher element in 36.3 and number in array is 37",
			in:   input{numbers: []int{37}, target: 33},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When target exists in numbers and there are no close values",
			in:   input{numbers: []int{1, 2, 3, 4, 5}, target: 3},
			out:  SearchResult{Index: 2, Number: 3},
			err:  false,
		},
		{
			name: "When target exists in numbers and there are close values",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 75, 80, 90, 100, 110, 120, 130}, target: 80},
			out:  SearchResult{Index: 8, Number: 80},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there are no close values",
			in:   input{numbers: []int{1, 2, 4, 5}, target: 3},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When target does not exists in numbers and there is single close values",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 75, 90, 100, 110, 120, 130}, target: 80},
			out:  SearchResult{Index: 7, Number: 75},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there are multiple close values, returns closest",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 73, 85, 90, 100, 110, 120, 130}, target: 80},
			out:  SearchResult{Index: 8, Number: 85},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there are multiple close values with the same distance to target, returns lower",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 75, 85, 90, 100, 110, 120, 130}, target: 80},
			out:  SearchResult{Index: 7, Number: 75},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there is one clost value is on index 0, returns it",
			in:   input{numbers: []int{12, 20, 30, 40, 50, 60, 70, 75, 85, 90, 100, 110, 120, 130}, target: 11},
			out:  SearchResult{Index: 0, Number: 12},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there is one close value on last index, returns it",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 75, 85, 90, 100, 110, 120, 130}, target: 131},
			out:  SearchResult{Index: 13, Number: 130},
			err:  false,
		},
		{
			name: "When target does not exists in numbers and there is closes value is on index 0 and is out of range",
			in:   input{numbers: []int{10, 20, 30, 40, 50, 60, 70, 75, 85, 90, 100, 110, 120, 130}, target: 1},
			out:  SearchResult{},
			err:  true,
		},
		{
			name: "When target does not exists in numbers and there is closest value which is out of range",
			in:   input{numbers: []int{10, 20}, target: 80},
			out:  SearchResult{},
			err:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotOut, gotErr := FindIndexOrClosest(test.in.numbers, test.in.target)

			if gotErr != nil && !test.err {
				t.Errorf("Unexpected error: %v", gotErr)
			}

			if test.err && gotErr == nil {
				t.Errorf("Expect error, but it was not returned")
			}

			if !reflect.DeepEqual(test.out, gotOut) {
				t.Errorf("Expected: %v \n Got: %v \n", test.out, gotOut)
			}
		})
	}
}
