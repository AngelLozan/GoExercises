package main

import "testing"
// @dev go test -cover gives test coverage. 
// @dev What kind of data I want to test.
// Define struct (top)
// Then every entry in the bottom (data) will be test.
var tests =[]struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false}, //each value corresponds to struct definition above
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}


func TestDivision(t *testing.T){
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one") //Fail test if error is expected tt.isErr (true or false)
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error()) //Fail test if error is not expected
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T){
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}


func TestBadDivide(t *testing.T){
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
