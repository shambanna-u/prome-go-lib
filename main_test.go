package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	var tests = []struct {
		input  string
		output int
	}{
		{"https://google.com", 200},
		{"https://gmail.com", 200},
		{"https://httpstat.us/200", 200},
		{"https://httpstat.us/503", 503},
	}
	for _, test := range tests {

		out := instrumentedHandler(test.input)
		if test.output != out {
			t.Fatal("Test Failed expected {} got {}", test.output, out)
		}

	}

}
