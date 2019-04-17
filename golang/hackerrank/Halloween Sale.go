package main

import "fmt"

// Complete the howManyGames function below.
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var result int32
	var count int32
	for result < s {
		if p-(count*d) < m {
			result += m
		} else {
			result += p - (count * d)
		}
		count++
	}
	if result == s {
		return count
	} else {
		return count - 1
	}

}

func main() {
	fmt.Println(howManyGames(1, 1, 1, 9981))
}
