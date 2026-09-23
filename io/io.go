package io

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)


func CheckDigits() (bool, int) {
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

func HandleCSVFile() [][]string {
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

func SaveToCSVFile(stat []string, csvFile [][]string) error {
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