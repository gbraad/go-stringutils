/*
Copyright 2017 Gerard Braad <me@gbraad.nl>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package strings

import (
	"reflect"
	"runtime"
	"testing"
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestHasMatcher(t *testing.T) {
	var testCases = []struct {
		testString     string
		matcher        func(string) bool
		expectedResult bool
	}{
		{"abc", HasLetters, true},
		{"123", HasNumbers, true},
		{"abc", HasNumbers, false},
		{"123", HasLetters, false},
		{"abc123", HasLetters, true},
		{"abc123", HasNumbers, true},
		{"abc", HasOnlyLetters, true},
		{"123", HasOnlyNumbers, true},
		{"abc123", HasOnlyLetters, false},
		{"abc123", HasOnlyNumbers, false},
		{"!@#$%^&*()", HasLetters, false},
		{"!@#$%^&*()", HasNumbers, false},
		{"!@#$%^&*()", HasOnlyLetters, false},
		{"!@#$%^&*()", HasOnlyNumbers, false},
	}

	for _, testCase := range testCases {
		actualResult := testCase.matcher(testCase.testString)
		if actualResult != testCase.expectedResult {
			t.Errorf("Unexpected result for '%s'. Expected '%s' to return '%t', but got '%t'", getFunctionName(testCase.matcher), testCase.testString, testCase.expectedResult, actualResult)
		}
	}
}

func TestGetMatcher(t *testing.T) {
	var testCases = []struct {
		testString     string
		matcher        func(string) string
		expectedResult string
	}{
		{"abc", GetOnlyLetters, "abc"},
		{"123", GetOnlyNumbers, "123"},
		{"abc123", GetOnlyLetters, "abc"},
		{"abc123", GetOnlyNumbers, "123"},
		{"!@#$%^&*()", GetOnlyLetters, ""},
		{"!@#$%^&*()", GetOnlyNumbers, ""},
		{"123GB", GetSignedNumbers, "123"},
		{"-123GB", GetSignedNumbers, "-123"},
	}

	for _, testCase := range testCases {
		actualResult := testCase.matcher(testCase.testString)
		if actualResult != testCase.expectedResult {
			t.Errorf("Unexpected result for '%s'. Expected '%s' to return '%t', but got '%t'", getFunctionName(testCase.matcher), testCase.testString, testCase.expectedResult, actualResult)
		}
	}
}

func TestStartsWithMatcher(t *testing.T) {
	var testCases = []struct {
		testString     string
		matcher        func(string) bool
		expectedResult bool
	}{
		{"abc", StartsWithLetter, true},
		{"abc", StartsWithNumber, false},
		{"123", StartsWithLetter, false},
		{"123", StartsWithNumber, true},
	}

	for _, testCase := range testCases {
		actualResult := testCase.matcher(testCase.testString)
		if actualResult != testCase.expectedResult {
			t.Errorf("Unexpected result for '%s'. Expected '%s' to return '%t', but got '%t'", getFunctionName(testCase.matcher), testCase.testString, testCase.expectedResult, actualResult)
		}
	}
}
