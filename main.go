package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	randomNum := rand.Intn(max-min+1) + min

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	fmt.Println(`Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)

	diff := 1
	fmt.Print("Enter your choice: ")
	fmt.Scan(&diff)

	if diff > 3 || diff < 1 {
		fmt.Println("Keuze niet goed, defaulting to Easy")
		diff = 1
	}

	chances := 10
	diffString := "Easy"

	switch diff {
	case 2:
		chances = 5
		diffString = "Medium"
	case 3:
		chances = 3
		diffString = "Hard"
	}

	fmt.Printf("Great! You have selected the %s difficulty level.\n", diffString)
	fmt.Printf("You have %d chances. Let's start the game!\n", chances)

	// Gebruik nu 'chances' in de loop
	for i := 0; i < chances; i++ {
		input := askInput()

		if input < randomNum {
			fmt.Println("HOGER!")
		} else if input > randomNum {
			fmt.Println("LAGER!")
		} else {
			fmt.Println("JUIST! HET NUMMER WAS", randomNum)
			return
		}

		remaining := chances - i - 1

		if remaining > 1 {
			fmt.Printf("Nog %d kansen\n", remaining)
		} else if remaining == 1 {
			fmt.Println("Nog 1 kans")
		} else {
			fmt.Println("Geen kansen meer! Het nummer was", randomNum)
		}
	}
}

func askInput() int {
	fmt.Print("Enter your choice: ")
	userInput := 0
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Er ging iets fout", err)
		return -1
	}

	return userInput
}