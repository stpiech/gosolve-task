package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectedVal int
		err         bool
	}{
		{
			name:        "Sets valid logLevel",
			in:          "Error",
			expectedVal: logLevelMap["Error"],
			err:         false,
		},
		{
			name:        "Sets not valid logLevel",
			in:          "Some random status",
			expectedVal: 0,
			err:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotErr := SetLogLevel(test.in)

			if gotErr != nil && !test.err {
				t.Errorf("Unexpected error: %v", gotErr)
			}

			if test.err && gotErr == nil {
				t.Errorf("Expect error, but it was not returned")
			}

			if test.expectedVal != LogLevel && gotErr == nil {
				t.Errorf("Expected: %v \n Got: %v \n", test.expectedVal, LogLevel)
			}
		})
	}
}

func TestDebugLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	debugLogger.SetOutput(buffer)

	LogLevel = logLevelMap["Debug"]
	DebugLogger("This is a debug message")
	if !strings.Contains(buffer.String(), "This is a debug message") {
		t.Errorf("Expected debug message not found in output")
	}
}

func TestInfoLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	infoLogger.SetOutput(buffer)

	LogLevel = logLevelMap["Info"]
	InfoLogger("This is a info message")
	if !strings.Contains(buffer.String(), "This is a info message") {
		t.Errorf("Expected debug message not found in output")
	}
}

func TestErrorLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	errorLogger.SetOutput(buffer)

	LogLevel = logLevelMap["Error"]
	ErrorLogger("This is a error message")
	if !strings.Contains(buffer.String(), "This is a error message") {
		t.Errorf("Expected error message not found in output")
	}
}
