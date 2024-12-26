package main

import (
	"fmt"
	"sort"
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

	// Mencari Ketua RT dan Wakil Ketua RT
	candidates := make([]int, 0, len(validVotes))
	for candidate := range validVotes {
		candidates = append(candidates, candidate)
	}

	sort.Slice(candidates, func(i, j int) bool {
		if validVotes[candidates[i]] == validVotes[candidates[j]] {
			return candidates[i] < candidates[j]
		}
		return validVotes[candidates[i]] > validVotes[candidates[j]]
	})

	ketuaRT := 0
	wakilRT := 0
	if len(candidates) > 0 {
		ketuaRT = candidates[0]
	}
	if len(candidates) > 1 {
		wakilRT = candidates[1]
	}

	// Output hasil
	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", totalValidVotes)

	for i := 1; i <= 20; i++ {
		if count, exists := validVotes[i]; exists {
			fmt.Printf("%d: %d\n", i, count)
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketuaRT)
	fmt.Printf("Wakil RT: %d\n", wakilRT)
}
