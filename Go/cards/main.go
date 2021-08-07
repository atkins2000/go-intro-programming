package main

//run two files go run main.go deck.go
func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" //equivalent statement as above := assign = change

	cards := newDeck()
	//cards = append(cards, "Six of Spades")

	//hand, remainingCards := deal(cards, 5)
	//cards.saveToFile("my_cards") //saves to file
	cards.shuffle()
	cards.print()
}
