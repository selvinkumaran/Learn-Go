package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(4, 4)
	if res != 8 {
		t.Errorf("Incorrect got %d, and want %d.", res, 8)
	}
}

// func TestAdd(t *testing.T) {
// 	got := add(4, 4)
// 	want := 8
// 	if got != want {
// 		log.Fatalf("Incorrect got %d, and want %d.", got, want)
// 	}
// }
