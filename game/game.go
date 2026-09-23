package game

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"slices"
	// "math/rand"
)

// Pick random word

// func randomSecretWord() string {
// 	content, err := os.ReadFile("wordle-words.txt")
// 	if err != nil {
// 		fmt.Println("Error when reading stat file:", err)
// 	}
// 	wordsArr := strings.Split(string(content), "\n")
// 	scWord := wordsArr[rand.Intn(len(wordsArr))]
// 	return scWord
// }



func PlayGame(scanner *bufio.Scanner, number int, username string) []string {
	wordle, wordlist := SecretWord(number)
	attempts := 1
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	failedString := ""
	correctMatch := false
	fmt.Print("Welcome to Wordle! Guess the 5-letter word.\n")
	GAMELOOP:
	for i:=5; i>=0; i--{
		fmt.Print("Enter your guess:")
		if scanner.Scan(){
			guess:= strings.TrimSpace(scanner.Text())
			isValid := IsGuessValid(guess, wordlist)
			if !isValid {
				i = i+1
			} else {
				failedString, correctMatch, alphabet = CheckWordle(guess, wordle, alphabet)
				if correctMatch {
					attempts = 6-i
					fmt.Println("Congratulations! You've guessd the word correctly")
					break GAMELOOP
				} else {
					fmt.Println(" Feedback: " + failedString)
					fmt.Print("Remaining letters:")
					for _, c :=range alphabet {
						fmt.Print(" "+c)
					}
					fmt.Println("\nAttempts remaining: ", i)
				}
				
			}
		} else {
			os.Exit(0)
		}
	}
	
	if correctMatch {
		return []string{username, wordle, strconv.Itoa(attempts), "win"}
	} else {
		fmt.Print("Game over. The correct word was: " + wordle)
		return []string{username, wordle, strconv.Itoa(attempts), "loss"}
	}
}

func SecretWord(number int) (string, []string) {
	content, err := os.ReadFile("wordle-words.txt")
	if err != nil {
		fmt.Println("Error when reading wordle file:", err)
	}
	wordleArr := strings.Split(string(content), "\n")
	return wordleArr[number], wordleArr
}

func IsGuessValid(input string, wordlist []string) (bool) {
	checkAlphabet := true
	isValidInput:= true
	if len(input) !=5 {
		fmt.Println(" Your guess must be exactly 5 letters long.")
		isValidInput = false
	} else if checkAlphabet {
		CHECKALPHABETLOOP:
		for _, r:=range input {
			if !(r >= 'a' && r <= 'z') {
				fmt.Println("Your guess must only contain lowercase letters.")
				isValidInput = false
				break CHECKALPHABETLOOP
			} 
		}
	} else if !slices.Contains(wordlist, input) {
		fmt.Println("Word not in list. Please enter a valid word.")
		isValidInput = false
	}
	return isValidInput
}

func CheckWordle(input string, wordle string, alphabet []string) (string, bool, []string) {
	Green := "\033[32m"
	Yellow := "\033[33m"
	Gray := "\033[37m"
	Reset := "\033[0m"

	correctMatch := false
	testStr := ""
	inputArr := strings.Split(input, "")
	wordleArr := strings.Split(wordle, "")
	if input == wordle {
		testStr = wordle
		correctMatch = true
	} else {
		for j:=0; j<len(inputArr); j++ {
			if strings.Contains(wordle, inputArr[j]) {
				if inputArr[j] == wordleArr[j] {
					testStr += Green + strings.ToUpper(inputArr[j]) + Reset
				} else {
					testStr += Yellow + strings.ToUpper(inputArr[j]) + Reset
				}
			} else {
				testStr += Gray + strings.ToUpper(inputArr[j]) + Reset
				indexInAlphabet := slices.Index(alphabet, strings.ToUpper(inputArr[j]))
				if indexInAlphabet >=0 {
					alphabet = append(alphabet[:indexInAlphabet], alphabet[indexInAlphabet+1:]...)
				}
			}
		}
	}
	return testStr, correctMatch, alphabet
}