package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)

	if result != "Foo" {
		t.Errorf("Incorrect. Received: %s Expected: %s ", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	// define columns of the table
	var tests = []struct {
		name 		string
		input		int
		expected	string
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
	var data = []struct {
		title 	string
		input	int
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
			t.Parallel()
			result := Fooer(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

// Todo: Skipping tests, but first Look at why warnings ^^