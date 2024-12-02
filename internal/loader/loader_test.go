package loader

import (
	"os"
	"reflect"
	"testing"
)

func TestFileToSlice(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		out         []int
		err         bool
	}{
		{
			name:        "When file is correct",
			fileContent: "1\n2\n3\n4\n5",
			out:         []int{1, 2, 3, 4, 5},
			err:         false,
		},
		{
			name:        "When file has element which can't be converted to int",
			fileContent: "1\nhello\n3\n4\n5",
			out:         []int{},
			err:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "testdata")

			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}

			defer os.Remove(tempFile.Name())

			tempFile.WriteString(test.fileContent)
			tempFile.Close()

			gotOut, gotErr := FileToSlice(tempFile.Name())

			if gotErr != nil && !test.err {
				t.Errorf("Unexpected error: %v", gotErr)
			}

			if test.err && gotErr == nil {
				t.Errorf("Expect error, but it was not returned")
			}

			if !reflect.DeepEqual(test.out, gotOut) && !test.err {
				t.Errorf("Expected: %v \n Got: %v \n", test.out, gotOut)
			}
		})
	}
}
