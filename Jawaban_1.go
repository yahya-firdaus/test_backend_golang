package main

import (
	"fmt"
)

// Fungsi untuk menghitung weight dari char
func charWeight(char rune) int {
	return int(char - 'a' + 1)
}

// Hitung weight dari string
func pengecekanWeight(s string, queries []int) []string {
	weights := make(map[int]bool)
	
	n := len(s)
	for i := 0; i < n; {
		// Inisiasi dengan char dan weight nya
		char := s[i]
		weight := charWeight(rune(char))
		
		// Menghitung char yang sama
		count := 1
		for j := i + 1; j < n && s[j] == char; j++ {
			count++
		}
		
		// Menambahkan weight
		for k := 1; k <= count; k++ {
			weights[weight*k] = true
		}
		
		// Next ke char selanjutnya
		i += count
	}

	results := make([]string, len(queries))

	for i, query := range queries {
		if weights[query] {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	return results
}

func main() {
	// Contoh 1
	s := "abbcccd"
	queries := []int{1, 3, 9, 8}
	
	// Menghitung weight
	result := pengecekanWeight(s, queries)
	
	// Print hasilnya
	fmt.Println(result)
}
