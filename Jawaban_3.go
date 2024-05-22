package main

import (
	"fmt"
	"strconv"
)

// Fugsi rekursif untuk highest palindrome
func mencariPalindromTertinggiRecursive(s []rune, left int, right int, k int) []rune {
	// Jika sebelah kiri lebih besar daripada kanan, maka akan return string
	if left > right {
		return s
	}

	// Jika char sudah sama
	if s[left] == s[right] {
		return mencariPalindromTertinggiRecursive(s, left+1, right-1, k)
	}

	// Pengecekan jika diharuskan me replace char
	if k > 0 {
		if s[left] > s[right] {
			s[right] = s[left]
		} else {
			s[left] = s[right]
		}
		return mencariPalindromTertinggiRecursive(s, left+1, right-1, k-1)
	}

	// Jika tidak ada yang perlu untuk di replace dan char tidak sama maka mengembalikan nil
	return nil
}

// Fungsi untuk maximize nilai palindrom
func maxPalindrom(s []rune, left int, right int, k int) []rune {
	if left > right {
		return s
	}

	if s[left] == s[right] {
		return maxPalindrom(s, left+1, right-1, k)
	}

	if k > 1 {
		if s[left] != '9' {
			s[left] = '9'
			k--
		}
		if s[right] != '9' {
			s[right] = '9'
			k--
		}
		return maxPalindrom(s, left+1, right-1, k)
	}

	return s
}

// Fungsi untuk menentukan palindrom yang tertinggi
func prosesPalindrome(s string, k int) string {
	// Mengubah string menjadi rune
	rs := []rune(s)
	n := len(rs)

	rs = mencariPalindromTertinggiRecursive(rs, 0, n-1, k)
	if rs == nil {
		return "-1"
	}

	rs = maxPalindrom(rs, 0, n-1, k)

	return string(rs)
}

func main() {
	// Contoh inputan
	input1 := "3943"
	k1 := 1
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input1, k1, prosesPalindrome(input1, k1))

	input2 := "932239"
	k2 := 2
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input2, k2, prosesPalindrome(input2, k2))

	input3 := "12321"
	k3 := 1
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input3, k3, prosesPalindrome(input3, k3))
}
