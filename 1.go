package main

import (
	"fmt"
)

func main() {
	var votes []int
	var input int

	// Membaca input sampai ditemukan angka 0
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		votes = append(votes, input)
	}

	totalVotes := len(votes)
	validVotes := make(map[int]int)

	// Memvalidasi suara
	for _, vote := range votes {
		if vote >= 1 && vote <= 20 {
			validVotes[vote]++
		}
	}

	totalValidVotes := 0
	for _, count := range validVotes {
		totalValidVotes += count
	}

	// Output hasil
	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", totalValidVotes)

	for i := 1; i <= 20; i++ {
		if count, exists := validVotes[i]; exists {
			fmt.Printf("%d: %d\n", i, count)
		}
	}
}
