package main

// go test in terminal
// in go, there is no testing framework like rspec, jest, enzyme
// all go knows is that we ran a function and there might be something wrong with it
// in go, test fileNames should be name_test.go by convention
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

  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card in the deck to be Ace of Spades, but got %v", d[0])
  }

  if d[len(d) - 1] != "King of Clubs" {
    t.Errorf("Expected last card in the deck to be King of Clubs, but got %v", d[len(d) - 1])
  }
}
