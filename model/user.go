package user

import (
	"strconv"
)

func CheckStats(filteredStatsArr [][]string , username string) (int, int, float64 ) {
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

