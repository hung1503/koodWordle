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

func main() {
	startTheGame, number := checkDigits()
	if startTheGame {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("----Welcome to Wordle----")
		fmt.Println("Enter your username:")
		scanner := bufio.NewScanner(os.Stdin)

		for {
		
			// Inputing username
			username := strings.TrimSpace(scanner.Text())
			USERNAMELOOP:
			for{
				if scanner.Scan(){
					username = strings.TrimSpace(scanner.Text())
					if len(username) == 0 {
					fmt.Println("Invalid username! Please try again")
					} else {
						break USERNAMELOOP
					}
				} else {
					fmt.Println("Exiting program...")
					os.Exit(0)
				}
			}
			
			fmt.Println("Hi", username)
			csvFile := handleCSVFile()
			matchStats := playGame(scanner, number, username)
			err := saveToCSVFile(matchStats, csvFile)
			csvFile = append(csvFile, matchStats)
			if err != nil {
				fmt.Println("Error with saving game stat in CSV file")
			}
			STATSLOOP:	
			for {
				fmt.Println("Do you want to see your stats? (yes/no)")
				if scanner.Scan() {
					answer := strings.TrimSpace(scanner.Text())
					if answer == "yes" || answer == "y" {
						// csvFile := handleCSVFile()
						gameCount, winCount, aveAttemps :=checkStats(csvFile, username)
						fmt.Println("Stat for", username)
						fmt.Println("Game played:", gameCount)
						fmt.Println("Game won:", winCount)
						fmt.Println("Average attempts per game:", aveAttemps)
						break STATSLOOP
					} else {
						break STATSLOOP
					}
				} else {
					fmt.Println("Exiting program...")
					os.Exit(0)
				}
			}
			fmt.Println("Press Enter to exit...")
			for {
				if scanner.Scan(){
					input := scanner.Text()
					if input == "" {
						os.Exit(0)
					} 
				} 
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

func handleCSVFile() [][]string {
	content, err := os.ReadFile("stats.csv")
	if err != nil {
		fmt.Println("Error when reading stat file:", err)
	}
	statsArr := strings.Split(string(content), "\n")
	filteredStatsArr := [][]string{}
	for _, value :=range statsArr {
		oneRow := strings.Split(value, ",")
		filteredStatsArr = append(filteredStatsArr, oneRow)
	}	
	return filteredStatsArr
}

func checkStats(filteredStatsArr [][]string , username string) (int, int, float64 ) {
	gameCount := 0
	winCount := 0
	var aveAttemps float64 = 0
	for _, stat := range filteredStatsArr {
		if stat[0] == username {
			gameCount++
			if stat[3] == "win"{
				winCount++
			}
			attempts, _:=strconv.ParseFloat(stat[2], 64)
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
	return wordleArr[number-1], wordleArr
}

func playGame(scanner *bufio.Scanner, number int, username string) []string {
	wordle, wordlist := secretWord(number)
	attempts := 1
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	failedString := ""
	correctMatch := false
	GAMELOOP:
	for i:=5; i>=0; i--{
		fmt.Print("Enter your guess. 5-LETTER word only: ")
		if scanner.Scan(){
			guess:= strings.TrimSpace(scanner.Text())
			isValid := isGuessValid(guess, wordlist)
			if !isValid {
				i = i+1
			} else {
				failedString, correctMatch, alphabet = checkWordle(guess, wordle, alphabet)
				if correctMatch {
					attempts = 6-i
					fmt.Println("Congratulations! You've guessd the word correctly")
					break GAMELOOP
				} else {
					fmt.Println("Feedback: " + failedString)
					fmt.Println("Remaining letters: ", alphabet)
					fmt.Println("Attemps remaining: ", i)
				}
				
			}
		}
	}
	fmt.Println("The wordle is " + wordle)
	if correctMatch {
		return []string{username, wordle, strconv.Itoa(attempts), "win"}
	} else {
		return []string{username, wordle, strconv.Itoa(attempts), "loss"}
	}
}

func isGuessValid(input string, wordlist []string) (bool) {
	checkAlphabet := true
	isValidInput:= true
	if len(input) !=5 {
		fmt.Println("The word must exactly 5 letters long")
		isValidInput = false
	} else if checkAlphabet {
		CHECKALPHABETLOOP:
		for _, r:=range input {
			if !(r >= 'a' && r <= 'z') {
				fmt.Println("The word must only contains lowercase letters")
				isValidInput = false
				break CHECKALPHABETLOOP
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
		for j:=0; j<len(inputArr); j++ {
			if strings.Contains(wordle, inputArr[j]) {
				if inputArr[j] == wordleArr[j] {
					testStr += Green + strings.ToUpper(inputArr[j]) + Reset
				} else {
					testStr += Yellow + strings.ToUpper(inputArr[j]) + Reset
				}
			} else {
				testStr += White + strings.ToUpper(inputArr[j]) + Reset
				indexInAlphabet := slices.Index(alphabet, strings.ToUpper(inputArr[j]))
				if indexInAlphabet >=0 {
					alphabet = append(alphabet[:indexInAlphabet], alphabet[indexInAlphabet+1:]...)
				}
			}
		}
	}
	return testStr, correctMatch, alphabet
}

func saveToCSVFile(stat []string, csvFile [][]string) error {
	f, err := os.OpenFile("stats.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content := strings.Join(stat, ",")
	 _, er := f.WriteString("\n"+content)
	 if er != nil {
		panic(err)
	}
	return er
}