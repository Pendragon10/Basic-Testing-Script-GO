package main

import "testing"

func TestSportRating(t *testing.T) {

	//Testing for one argument
	fSport := SportRating("Football")

	if fSport != "Football: 100/10" {
		t.Errorf("Failed: Expected %v, got %v", "Football: 100/10", fSport)
	} else {
		t.Logf("Success: Expected %v, got %v", "Football: 100/10", fSport)
	}
}

func TableDrivenTestSportsRating(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"soccer", "Soccer: 9/10"},
		{"pingpong", "Pingpong: 10/10"},
		{"", "Invalid Sport or Argument"},
	}

	for _, test := range tests {
		if output := SportRating(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received", test.input, test.expected, output)
		}
	}

}
