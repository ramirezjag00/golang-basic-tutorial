package main

import "fmt"

func main() {
  // cards := newDeck()

  // variable using print method for type deck
  // cards.print()

  // hands, remainingCards := deal(cards, 5)
  //
  // hands.print()
  // remainingCards.print()

  greeting := "Hello there!"
  fmt.Println([]byte(greeting))
  // this is the result of "type conversion"
  // basically turning string characters into its value in ascii (byte) and putting results in slice
  // [72 101 108 108 111 32 116 104 101 114 101 33]
}
