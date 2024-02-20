package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := Paradise(" Australia")
	want := "My place is Australia"

	if got != want {
		log.Fatalf("Error -got %s and want %s", got, want)
	}
}
