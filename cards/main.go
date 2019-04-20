package main

func main() {
  // cards := newDeck()
  // cards.saveToFile("my_cards")

  // if you run this once, will create a file in the same directory, if you edit/save the file, then run this package again, it will overwrite your changes in the file

  cards := newDeckFromFile("my_cards")
  cards.print()

  // will produce error if fileName inputted is non-existent
}
