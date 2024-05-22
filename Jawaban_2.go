package main

import (
	"fmt"
)

// Fungsi untuk melakukan cek jika simbol balance
func pengecekanSimbol(s string) string {
	// Stack untuk mengecek simbol
	tempStack := []rune{}

	// Untuk mencocokkan simbol pembuka dan penutup
	simbol := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		// Jika simbol pembuka maka akan masuk kedalam stack
		case '(', '[', '{':
			tempStack = append(tempStack, char)
		// Jika simbol penutup maka akan dilakukan pengecekan
		case ')', ']', '}':
			// Jika stack kosong atau simbol teratas di stack tidak cocok maka akan return "NO"
			if len(tempStack) == 0 || tempStack[len(tempStack)-1] != simbol[char] {
				return "NO"
			}
			// Hapus simbol paling atas dari stack
			tempStack = tempStack[:len(tempStack)-1]
		}
	}

	// Jika stack kosong, semua simbol balance
	if len(tempStack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	// Contoh input
	inputs := []string{
		"{ [ ( ) ] }",
		"{ [ ( ] ) }",
		"{ ( ( [ ] ) [ ] ) [ ] }",
	}

	// Pengecekan dari input dan print hasilnya
	for _, input := range inputs {
		result := pengecekanSimbol(input)
		fmt.Printf("Input: %s\nOutput: %s\n", input, result)
	}
}
