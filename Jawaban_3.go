package main

import (
	"fmt"
	"strconv"
)

// Fugsi rekursif untuk highest palindrome
func highestPalindromeRecursive(s []rune, left int, right int, k int) []rune {
	// Jika sebelah kiri lebih besar daripada kanan, maka akan return string
	if left > right {
		return s
	}

	// Jika char sudah sama
	if s[left] == s[right] {
		return highestPalindromeRecursive(s, left+1, right-1, k)
	}

	// Pengecekan jika diharuskan me replace char
	if k > 0 {
		if s[left] > s[right] {
			s[right] = s[left]
		} else {
			s[left] = s[right]
		}
		return highestPalindromeRecursive(s, left+1, right-1, k-1)
	}

	// Jika tidak ada yang perlu untuk di replace dan char tidak sama maka mengembalikan nil
	return nil
}

// Fungsi untuk maximize nilai palindrom
func maximizePalindrome(s []rune, left int, right int, k int) []rune {
	if left > right {
		return s
	}

	if s[left] == s[right] {
		return maximizePalindrome(s, left+1, right-1, k)
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
		return maximizePalindrome(s, left+1, right-1, k)
	}

	return s
}

// Fungsi untuk menentukan palindrom yang tertinggi
func highestPalindrome(s string, k int) string {
	// Mengubah string menjadi rune
	rs := []rune(s)
	n := len(rs)

	rs = highestPalindromeRecursive(rs, 0, n-1, k)
	if rs == nil {
		return "-1"
	}

	rs = maximizePalindrome(rs, 0, n-1, k)

	return string(rs)
}

func main() {
	// Contoh inputan
	input1 := "3943"
	k1 := 1
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input1, k1, highestPalindrome(input1, k1))

	input2 := "932239"
	k2 := 2
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input2, k2, highestPalindrome(input2, k2))

	input3 := "12321"
	k3 := 1
	fmt.Printf("Input: %s, k: %d\nOutput: %s\n", input3, k3, highestPalindrome(input3, k3))
}
