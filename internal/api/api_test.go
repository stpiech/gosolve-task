package api

import (
	"testing"
)

func TestTransformValueParam(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  int
		err  bool
	}{
		{
			name: "When valid stringified integer is passed",
			in:   "5",
			out:  5,
			err:  false,
		},
		{
			name: "When value is not provided",
			in:   "",
			out:  0,
			err:  true,
		},
		{
			name: "When value is not atoiable",
			in:   "hello",
			out:  0,
			err:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotOut, gotErr := transformValueParam(test.in)

			if gotErr != nil && !test.err {
				t.Errorf("Unexpected error: %v", gotErr)
			}

			if test.err && gotErr == nil {
				t.Errorf("Expect error, but it was not returned")
			}

			if test.out != gotOut {
				t.Errorf("Expected: %v \n Got: %v \n", test.out, gotOut)
			}
		})
	}

}
