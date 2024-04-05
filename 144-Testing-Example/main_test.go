package main

import (
	"testing"
)

func TestAddSum(t *testing.T) {
	result := AddSum(2, 3)
	if result != 5 {
		t.Error("Expected", 5, "got", result)
	}

}
