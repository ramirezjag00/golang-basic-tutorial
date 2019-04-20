package main

import (
	"fmt"
  // ioutil is a subpackage of io
  "io/ioutil"
  "os"
	"strings"
)

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

// create a helper method called toString which returns one large string that came from joined strings from the deck type of cards
func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

// a helper method for using ioutil that writes/creates data/file in our device
// it accepts fileName of string that uses type of deck data, and may/may not return a type of error
// WriteFile method accepts a fileName, byte slice using type conversion, and 0666 which is the OS permission 0666 means anyone can able to read write and edit the file
func (d deck) saveToFile(fileName string) error {
  return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// a helper method that accepts a fileName (reads fileName) of an existing file and returns a new deck type of cards
func newDeckFromFile(fileName string) deck {
  // read the fileName inputted and print if there is an error and quit
  bs, err := ioutil.ReadFile(fileName)
  if err != nil {
    // Option#1 - log error and return a call to newDeck()
    // Option#2 - log error and entirely quit the program
    // log.Fatal(err)
    // same as
    fmt.Println("ERROR: ", err)
    // non-zero value for os.Exit means unsuccessful, will exit the program
    os.Exit(1)
  }
  // if no error, convert the byte slice into one large string and comma split everything to make and return a new deck
  s := strings.Split(string(bs), ",")
  return deck(s)
}
