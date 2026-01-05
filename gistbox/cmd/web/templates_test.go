package main

import (
	"testing"
	"time"

	"snippetbox-webapp/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to the humanDate() function (the tm field), and expected output
	// (the expected field).
	tests := []struct {
		name	string
		tm		time.Time
		exp		string
	}{
		{
			name:	"UTC",
			tm:		time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			exp:	"17 Mar 2024 at 10:15",
		},
		{
			name:	"Empty",
			tm:		time.Time{},
			exp:	"",
		},
		{
			name:	"CET",
			tm:		time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			exp:	"17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		// Use t.Run() to run a subtest for each test case.
		// t.Run takes as first parameter the name of the test (used to 
		// identify the subtest in a log output) and the second parameter
		// is an anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			// Use the new assert.Equal() helperto compare the expected
			// and actual values.
			assert.Equal(t, hd, tt.exp)
		})
	}
}