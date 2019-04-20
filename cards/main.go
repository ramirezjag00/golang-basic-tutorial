package main

func main() {
  cards := newDeck()

  // variable using print method for type deck
  // cards.print()

  hands, remainingCards := deal(cards, 5)

  hands.print()
  remainingCards.print()
}
