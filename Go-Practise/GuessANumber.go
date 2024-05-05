package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	//Here 1 is added to number since rand.Int will return between 0-99
	fmt.Println("I have choosen a number between 1 to 100 ", target)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You have 10 chances to guess the number")
	success := false
	for i := 1; i <= 10; i++ {
		fmt.Println("Enter ", i, " number=>")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			i--
			continue
		}

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Your guess is less then target")
		} else if guess > target {
			fmt.Println("Your guess is greater then target")
		} else {
			success = true
			fmt.Println("Your guess is right,target is: ", target)
			break
		}
	}
	if !success {
		fmt.Println("You all guesses were wrong, Target is: ", target)
	}
}
