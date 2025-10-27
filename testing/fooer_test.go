package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Guide about testing: https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/

func TestFooer(t *testing.T) {
	result := Fooer(3)

	if result != "Foo" {
		t.Errorf("Incorrect. Received: %s Expected: %s ", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	// define columns of the table
	var tests = []struct {
		name     string
		input    int
		expected string
	}{
		// the table itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	// execute test table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fooer(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestFooerStopExecution(t *testing.T) {
	input := 3

	result := Fooer(3)

	t.Logf("The input was %d", input)

	if result != "Foo" {
		t.Errorf("Got %s; Want %s", result, "Foo")
	}

	// t.Fatalf("Stop the test now, we have seen enough")
	// t.Error("This won't be executed")
}

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel() // tell go to run this test in parallel

		if result := Fooer(3); result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})

	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()

		if result := Fooer(7); result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

func TestFooerParallelTable(t *testing.T) {
	t.Parallel()

	var data = []struct {
		title    string
		input    int
		expected string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	// Run table tests in parallel
	for _, tt := range data {
		t.Run(tt.title, func(t *testing.T) {
			t.Logf("Foer test: %q for %v", tt.title, tt.input)
			result := Fooer(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

// `go test -v -test.short` will skip this test
// If have mix of integration and unit, and just want to run unit
// add this to unit tests
func TestFooerSkipped(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	if result := Fooer(3); result != "Foo" {
		t.Errorf("Error. Wanted %s, Got %s", "Foo", result)
	}
}

// Read about test cleanup: https://ieftimov.com/posts/testing-in-go-clean-tests-using-t-cleanup/
func TestWithCleanup(t *testing.T) {
	if result := Fooer(3); result != "Foo" {
		t.Errorf("Error. Wanted %s, Got %s", "Foo", result)
	}

	t.Cleanup(func() {
		// do your cleanup logic here
		t.Log("Cleaning up test.")
	})
}

// Helper: improve logs when tests fail
func TestWithHelper(t *testing.T) {
	t.Helper()
	if result := Fooer(3); result != "Bar" {
		t.Errorf("Error. Wanted %s, Got %s", "Bar", result)
	}
}

// temp dir. Creates temporary dir for test and deletes folder when test has finished
func TestWithTempDir(t *testing.T) {
	temp := t.TempDir()
	fmt.Println(temp)
}

// with coverage: run tests with: `go test -cover`

// Benchmark test:
// Your test function needs to be in a *_test file.
// The name of the function must start with Benchmark.
// The function must accept *testing.B as the unique parameter.
// The test function must contain a for loop using b.N as its upper bound
// See more: https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/#writing-benchmark-tests
func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fooer(i)
	}
}

// Fuzz testing: Random input to discover bugs and edge cases
// Your test function needs to be in a _test file.
// The name of the function must start with Fuzz.
// The test function must accept testing.F as the unique parameter.
// The test function must define initial values, called seed corpus, with the f.Add() method.
// The test function must define a fuzz target
// See more: https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/#writing-fuzz-tests
func FuzzFooer(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		Fooer(a)
	})
}

// testify package: https://github.com/stretchr/testify
// run go get github.com/stretchr/testify
func TestMapWithTestify(t *testing.T) {
	// Assert equal
	assert.Equal(t, "Foo", Fooer(0), "0 is divisibly by 3, shour return Foo")

	// Assert Not equal
	assert.NotEqual(t, "Foo", Fooer(1), "1 is not divisible by 3, should not return Foo")

	// require (stops execution if test failure)
	require.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "3"})

	// assert equality
	assert.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2"})
}
