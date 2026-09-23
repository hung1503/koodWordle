package io

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)


func CheckDigits() (bool, int) {
	arguments := os.Args
	if len(os.Args) == 1 {
		fmt.Println("Please provide a number as command line argument")
		return false, -1
	} else if len(arguments) > 2 {
		fmt.Println("Invalid input! Enter one number")
		return false, -1
	}
	number, err := strconv.Atoi(string(arguments[1]))
	if err!=nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return false, -1
	}
	
	return true, number
}

func HandleCSVFile() [][]string {
	content, _ := os.ReadFile("stats.csv")
	statsArr := strings.Split(string(content), "\n")
	filteredStatsArr := [][]string{}
	for _, value :=range statsArr {
		oneRow := strings.Split(value, ",")
		filteredStatsArr = append(filteredStatsArr, oneRow)
	}	
	return filteredStatsArr
}

func SaveToCSVFile(stat []string, csvFile [][]string) error {
	f, err := os.OpenFile("stats.csv", os.O_APPEND|os.O_WRONLY, 0644)
	 if err != nil {
		panic(err)
	}
	defer f.Close()
	content := strings.Join(stat, ",")
	 _, er := f.WriteString("\n"+content)
	if er != nil {
		panic(er)
	}
	return er
}