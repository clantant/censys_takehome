package scan

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConnString(t *testing.T) {
	testCases := []struct {
		name           string
		testIP         string
		testPort       string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "BuildSuccess",
			testIP:         "127.0.0.1",
			testPort:       "3306",
			expectedOutput: "tcp(127.0.0.1:3306)/",
			expectedError:  nil,
		},
		{
			name:           "EmptyFail",
			testIP:         "",
			testPort:       "",
			expectedOutput: "",
			expectedError:  errors.New("Bad IP address format"),
		},
		{
			name:          "BadIP",
			testIP:        "someotherstring",
			testPort:      "3306",
			expectedError: errors.New("Bad IP address format"),
		},
		{
			name:          "BadPort",
			testIP:        "127.0.0.1",
			testPort:      "someotherstring",
			expectedError: errors.New("Bad Port format"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			actualOutput, err := buildConnString(tt.testIP, tt.testPort)
			a.Equal(tt.expectedError, err)
			a.Equal(tt.expectedOutput, actualOutput)
		})
	}
}
