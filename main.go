package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"math/rand"
)

type Stats struct {
	Username string
	SecretWord string
	NumOfAttempts string
	WinLose string
}

func main() {
	startTheGame := checkDigits()
	if startTheGame {
		fmt.Println("----Welcome to Wordle----")
		fmt.Println("Enter your username:")
		scanner := bufio.NewScanner(os.Stdin)

		for {
			if scanner.Scan() {
				// Inputing username
				username := strings.TrimSpace(scanner.Text())
				fmt.Println("Hi ", username)
				csvFile := handleCSVFile()
				playGame(scanner)		
				
				
				
				// Checking stats
				gameCount, winCount, aveAttemps :=checkStats(csvFile, username)
				fmt.Println(gameCount, winCount, aveAttemps)
			} else {
				fmt.Println("Exiting program...")
				os.Exit(0)
			}
		}
	}
}

func checkDigits() bool {
	digitsWord := os.Args[1]
	if digitsWord != "5" {
		fmt.Println("Invalid index of the word. The index should be 5")
		return false
	} 
	return true
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

func secretWord() string {
	content, err := os.ReadFile("wordle-words.txt")
	if err != nil {
		fmt.Println("Error when reading stat file:", err)
	}
	wordsArr := strings.Split(string(content), "\n")
	scWord := wordsArr[rand.Intn(len(wordsArr))]
	return scWord
}

func playGame(scanner *bufio.Scanner) {
	wordle := secretWord()
	attempts := 6
	fmt.Print("Enter your guess. 5-LETTER word only: ")
		for i:=attempts; i>0; i-- {
			if scanner.Scan(){
			
			guess:= strings.TrimSpace(scanner.Text())
			fmt.Println("Your guess ", guess, wordle)
		}
	}
}

func checkGuess(input string) string {
	if input
}