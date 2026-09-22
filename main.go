package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"slices"
	// "math/rand"
)

type Stats struct {
	Username string
	SecretWord string
	NumOfAttempts string
	WinLose string
}

func main() {
	startTheGame, number := checkDigits()
	if startTheGame {
		fmt.Println("----Welcome to Wordle----")
		fmt.Println("Enter your username:")
		scanner := bufio.NewScanner(os.Stdin)

		for {
			if scanner.Scan() {
				// Inputing username
				username := strings.TrimSpace(scanner.Text())
				if len(username) == 0 {
					fmt.Println("Invalid username! Please try again")
					if scanner.Scan() {

					}
				} 
				fmt.Println("Hi", username)
				
				playGame(scanner, number)
				STATSLOOP		
				for {
					fmt.Println("Do you want to see your stats? (yes/no)")
					if scanner.Scan() {
						answer := strings.TrimSpace(scanner.Text())
						if answer == "yes" || answer == 'y' {
							csvFile := handleCSVFile()
							gameCount, winCount, aveAttemps :=checkStats(csvFile, username)
							fmt.Println("Stat for " + username)
							fmt.Println("Game played: " + gameCount)
							fmt.Println("Game won: " + winCount)
							fmt.Println("Average attempts per game: " + aveAttemps)
						} else {
							break STATSLOOP
						}
					} else {
						fmt.Println("Exiting program...")
						os.Exit(0)
					}
				}
			
			} else {
				fmt.Println("Exiting program...")
				os.Exit(0)
			}
		}
	}
}

func checkDigits() (bool, int) {
	isPass := true
	arguments := os.Args[1]
	number, err := strconv.Atoi(string(arguments))
	if err!=nil {
		fmt.Println("Invalid input! Must be an non negative integer number")
		isPass = false
	} else if len(arguments) > 2 {
		fmt.Println("Invalid input! Enter one number")
		isPass = false
	} else if number > 14854 {
		fmt.Println("Invalid number! Enter smaller number")
		isPass = false
	}
	return isPass, number
}

func handleCSVFile() []Stats {
	content, err := os.ReadFile("stats.csv")
	if err != nil {
		fmt.Println("Error when reading stat file:", err)
	}
	statsArr := strings.Split(string(content), "\n")
	filteredStatsArr := []Stats{}
	for _, value :=range statsArr {
		oneRow := strings.Split(value, ",")
		stat := Stats{oneRow[0], oneRow[1], oneRow[2], oneRow[3]}
		filteredStatsArr = append(filteredStatsArr, stat)
	}	
	return filteredStatsArr
}

func checkStats(filteredStatsArr []Stats, username string) (int, int, float64 ) {
	gameCount := 0
	winCount := 0
	var aveAttemps float64 = 0
	for _, stat := range filteredStatsArr {
		if stat.Username == username {
			gameCount++
			if stat.WinLose == "win"{
				winCount++
			}
			attempts, _:=strconv.ParseFloat(stat.NumOfAttempts, 64)
			aveAttemps+=attempts
		}
	} 
	aveAttemps = aveAttemps/float64(gameCount)
	return gameCount, winCount, aveAttemps
}


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


func secretWord(number int) (string, []string) {
	content, err := os.ReadFile("wordle-words.txt")
	if err != nil {
		fmt.Println("Error when reading wordle file:", err)
	}
	wordleArr := strings.Split(string(content), "\n")
	return wordleArr[number + 1], wordleArr
}

func playGame(scanner *bufio.Scanner, number int) {
	wordle, wordlist := secretWord(number)
	attempts := 6
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	failedString := ""
	correctMatch := false
	fmt.Print("Enter your guess. 5-LETTER word only: ")
		GAMELOOP:
		for i:=attempts; i>0; i--{
			if scanner.Scan(){
				fmt.Print("Enter your guess. 5-LETTER word only: ")
				guess:= strings.TrimSpace(scanner.Text())
				isValid := isGuessValid(guess, wordlist)
				if !isValid {
					i = i+1
				} else {
					failedString, correctMatch, alphabet = checkWordle(guess, wordle, alphabet)
					if correctMatch {
						fmt.Println("Congratulations! You've guessd the word correctly")
						break GAMELOOP
					} else {
						fmt.Println("Feedback: " + failedString)
						fmt.Println("Remaining letters: ", alphabet)
						fmt.Println("Attemps remaining: ", (i-1))
					}
					
				}
		}
	}
}

func isGuessValid(input string, wordlist []string) (bool) {
	checkAlphabet := true
	isValidInput:= true
	if len(input) !=5 {
		fmt.Println("The word must exactly 5 letters long")
		isValidInput = false
	} else if checkAlphabet {
		for _, r:=range input {
			if !(r >= 'a' && r <= 'z') {
				fmt.Println("The word must only contains lowercase letters")
				isValidInput = false
			} 
		}
	} else if slices.Contains(wordlist, input) {
		fmt.Println("Word is not in the list. Please enter a valid word")
		isValidInput = false
	}
	return isValidInput
}

func checkWordle(input string, wordle string, alphabet []string) (string, bool, []string) {
	Green := "\033[32m"
	Yellow := "\033[33m"
	White := "\033[97m"
	Reset := "\033[0m"

	correctMatch := false
	testStr := ""
	inputArr := strings.Split(input, "")
	wordleArr := strings.Split(wordle, "")
	if input == wordle {
		testStr = wordle
		correctMatch = true
		fmt.Println()
	} else {
		for i:=0; i<=len(wordleArr); i++ {
			for j:=0; j<=len(inputArr); j++ {
				if strings.Contains(wordle, inputArr[j]) {
					if inputArr[j] == wordleArr[j] {
						testStr += Green + strings.toUpper(inputArr[j]) + Reset
					} else {
						testStr += Yellow + strings.toUpper(inputArr[j]) + Reset
					}
				} else {
					testStr += White + strings.toUpper(inputArr[j]) + Reset
					indexInAlphabet := slice.Index(alphabet, strings.toUpper(inputArr[j]))
					if indexInAlphabet >=0 {
						alphabet = append(alphabet[:indexInAlphabet], alphabet[indexInAlphabet+1:]...)
					}
				}
			}
		}
	}
	return testStr, correctMatch, alphabet
}