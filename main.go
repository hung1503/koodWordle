package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				fmt.Println("Enter your username:")
				input := strings.TrimSpace(scanner.Text())
				fmt.Println(input)
				
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
		fmt.Println(oneRow)
	}	
	return filteredStatsArr
}