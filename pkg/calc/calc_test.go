package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedResult float64
		err            error
	}{
		{
			name:           "simple",
			expression:     "2+1",
			expectedResult: 3,
			err:            nil,
		},
		{
			name:           "bracket",
			expression:     "2*(2+2)",
			expectedResult: 8,
			err:            nil,
		},
		{
			name:           "brackets",
			expression:     "(10+1)*(2+2)",
			expectedResult: 44,
			err:            nil,
		},
		{
			name:           "mult",
			expression:     "2+2*2",
			expectedResult: 6,
			err:            nil,
		},
		{
			name:           "div",
			expression:     "1/2",
			expectedResult: 0.5,
			err:            nil,
		},
		{
			name:           "div zero",
			expression:     "8/0",
			expectedResult: 0,
			err:            ErrDivisionByZero,
		},

		{
			name:           "hard",
			expression:     "1+32",
			expectedResult: 0,
			err:            ErrSomethingWentWrong,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}
}
