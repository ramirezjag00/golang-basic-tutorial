package main

// go test in terminal

import "testing"

// UpperCase test func to make it distinguishable from normal methods
// type *testing.T of the received t argument
func TestNewDeck(t *testing.T) {
  // d is instance or receiver of newDeck
  d := newDeck()

  // assume a length, 52 and show an error if not true
  if len(d) != 52 {
    // formatted error
    // %v represents len(d)
    t.Errorf("Expected deck length of 52, but got %v", len(d))
  }
}
