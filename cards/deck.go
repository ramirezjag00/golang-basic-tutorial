package main

import "fmt"

// create a new type of of 'deck'
// which a slice of strings
// there is no class in golang
type deck []string

// this will create a new deck type of cards that will mock all values and suits
func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

  // iterete through the suits and values to mock the whole deck of cards
  // first _ is saying that we are not using the index i so we replace it with an _
  // second _ is saying that we are not using the index j so we replace it with an _
  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, value+" of "+suit)
    }
  }

  return cards
}

// any variable of type deck now gets access to the "print" method
// d is the 'receiver' values of the variable used or actual copy of the assigned deck
// there is no this in golang
// by convention we use 1-3 letter abbreviation of the type, but you can always change it to anything
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

// deal method has 2 arguments which are type deck and type int, which returns multiple values (with specified types) which are both type deck. YES, returning multiple values in golang is supported.
func deal(d deck, handSize int) (deck, deck) {
  // return instance of d from index 0 to notIncluding handSize and instance of d from handSize to last index
  return d[:handSize], d[handSize:]
}
