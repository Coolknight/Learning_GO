package main

import (
	"fmt"
	"math/rand"
)

const gameSize = 6
const health = 10

// Knight represents a player in the game with a unique number and points.
type Knight struct {
	number int
	points int
	next   *Knight
}

// LinkedList represents a circular linked list of Knights.
type LinkedList struct {
	head *Knight
}

// new initializes a circular linked list of Knights with default points.
func (list *LinkedList) new() {
	var current *Knight
	for i := 1; i <= gameSize; i++ {
		newKnight := &Knight{number: i, points: health}
		if list.head == nil {
			list.head = newKnight
		} else {
			current.next = newKnight
		}
		current = newKnight
	}
	current.next = list.head
}

func main() {
	// Initialize the game with a circular linked list of Knights.
	game := LinkedList{}
	game.new()
	current := game.head

	// Continue the game until only one Knight is remaining.
	for current != current.next {
		// Simulate a hit by generating a random number and reducing the next Knight's points.
		rnd := rand.Intn(6) + 1
		current.next.points -= rnd

		// Display the hit information.
		fmt.Printf("K%d hits K%d for %d\n", current.number, current.next.number, rnd)

		// Check if the next Knight has run out of points and remove them from the list if necessary.
		if current.next.points < 1 {
			fmt.Printf("K%d dies\n", current.next.number)
			current.next = current.next.next
		}

		// Move to the next Knight in the circular linked list.
		current = current.next
	}

	// Display the winner.
	fmt.Printf("K%d wins\n", current.number)
}
