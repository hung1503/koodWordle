package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"koodWordle/game"
	"koodWordle/io"
	"koodWordle/model"
)

func main() {
	startTheGame, number := io.CheckDigits()
	if startTheGame {
		// os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Print("Enter your username: ")
		scanner := bufio.NewScanner(os.Stdin)
		OUTERLOOP:
		for {
			gameStart:=true
			enterUsername := true
			username:=""
			// Inputing username
			if number > 14854 || number < 0 {
				fmt.Print("Invalid word number.\n")
				gameStart=false
				enterUsername= false
				break OUTERLOOP
			}
			if enterUsername {
				USERNAMELOOP:
				for{
					if scanner.Scan(){
						username := strings.TrimSpace(scanner.Text())
						if len(username) == 0 {
						fmt.Println("Invalid username! Please try again")
						} else {
							break USERNAMELOOP
						}
					} else {
						gameStart=false
						break USERNAMELOOP
					}
				}
			}
			if gameStart {	
				csvFile := io.HandleCSVFile()
				matchStats := game.PlayGame(scanner, number, username)
				err := io.SaveToCSVFile(matchStats, csvFile)
				csvFile = append(csvFile, matchStats)
				if err != nil {
					fmt.Println("Error with saving game stat in CSV file")
				}
			
					fmt.Print("Do you want to see your stats? (yes/no):")
					if scanner.Scan() {
						answer := strings.TrimSpace(scanner.Text())
						if answer == "yes" || answer == "y" {
							// csvFile := handleCSVFile()
							gameCount, winCount, aveAttemps :=user.CheckStats(csvFile, username)
							fmt.Println("Stat for " + username + ":" )
							fmt.Println("Game played:", gameCount)
							fmt.Println("Game won:", winCount)
							fmt.Println("Average attempts per game:", aveAttemps)
							break OUTERLOOP
						} else {
							break OUTERLOOP
						}
					} else {
						fmt.Println("Exiting program...")
						os.Exit(0)
					}
				
			}
			
		} 
		fmt.Println("Press Enter to exit...")
		return
		// for {
		// 	if scanner.Scan(){
		// 		input := scanner.Text()
		// 		if input == ""  || input == "\n"{
		// 			os.Exit(0)
		// 		} 
		// 	} 
		// }
	}
}

