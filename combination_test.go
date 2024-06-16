package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, &ExampleTestSuite{})
}

func TestCombinationMaker(t *testing.T) {
	asserts := assert.New(t)
	testCases := []struct {
		testName string
		input    []string
		expected [][]string
		err      error
	}{
		{
			testName: "Single input returns single value",
			input: []string{
				"one",
			},
			expected: [][]string{
				{"one"},
			},
			err: nil,
		},
		//{
		//	testName: "Double input returns double value",
		//	input: []string{
		//		"one",
		//		"two",
		//	},
		//	expected: [][]string{
		//		{"one", "two"},
		//		{"two", "one"},
		//	},
		//	err: nil,
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			audit, err := CreateCombinations(tc.input)
			asserts.Equal(tc.expected, audit)
			asserts.Equal(tc.err, err)
		})
	}
}

func TestRemoveElementFunction(t *testing.T) {
	asserts := assert.New(t)
	testCases := []struct {
		testName  string
		input     []string
		index     int
		expOutput []string
		err       error
	}{
		{
			testName: "Ensure func removes middle element",
			input: []string{
				"zero",
				"one",
				"two",
			},
			index: 1,
			expOutput: []string{
				"zero",
				"two",
			},
			err: nil,
		},
		{
			testName: "Ensure func removes last element",
			input: []string{
				"zero",
				"one",
				"two",
			},
			index: 2,
			expOutput: []string{
				"zero",
				"one",
			},
			err: nil,
		},
		{
			testName: "Ensure func removes first element",
			input: []string{
				"zero",
				"one",
				"two",
			},
			index: 0,
			expOutput: []string{
				"one",
				"two",
			},
			err: nil,
		},
		{
			testName:  "Ensure empty array returns error",
			input:     []string{},
			index:     0,
			expOutput: []string(nil),
			err:       errors.New("Cannot split an empty array"),
		},
		{
			testName: "Ensure an array of a single item returns an error",
			input: []string{
				"Zero",
			},
			index:     0,
			expOutput: []string(nil),
			err:       errors.New("Cannot split an array with a single item"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			audit, err := RemoveElement(tc.input, tc.index)
			asserts.Equal(tc.expOutput, audit)
			asserts.Equal(tc.err, err)
		})
	}
}
