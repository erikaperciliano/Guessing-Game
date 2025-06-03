package game

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func MyLogic() {
	fmt.Println("Guessing Game")
	fmt.Println("A random number will be drawn. Try to guess it. The number is an integer between 0 and 100.")

	x := rand.Int64N(101)
	x = 10
	scanner := bufio.NewScanner(os.Stdin)
	guess := [10]int64{}

	for i := range guess {
		fmt.Println("What's your guess?")
		scanner.Scan()
		kick := scanner.Text()
		kick = strings.TrimSpace(kick) //removes all blank characters, spaces from the string...

		kickInt, err := strconv.ParseInt(kick, 10, 64)
		if err != nil {
			fmt.Println("Your guess needs to be a whole number!")
			return
		}

		switch {
		case kickInt < x:
			fmt.Println("You are wrong. The number drawn is greater than: ", kickInt)
		case kickInt > x:
			fmt.Println("You are wrong. The number drawn is less than:: ", kickInt)
		case kickInt == x:
			fmt.Printf(
				"Congratulations! You got it right! The number was: %d\n"+
					"You got it right in %d attempts.\n"+
					"These were his attempts: %v\n",
				x, i+1, guess[:i],
			)

			return
		}

		guess[i] = kickInt
	}

	fmt.Printf(
		"Unfortunately, you didn't get the number right, which was: %d\n"+
			"You had 10 attempts.\n"+
			"These were his attempts: %v\n",
		x, guess,
	)
}
