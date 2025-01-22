package main

import (
	"fmt"
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	grid := make([]string, r)
	for i := 0; i < r; i++ {
		fmt.Scan(&grid[i])
	}
	min_word := ""

	updateMinWord := func(word string) {
		if len(word) >= 2 {
			if min_word == "" || word < min_word {
				min_word = word
			}
		}
	}

	for _, row := range grid {
		current_word := ""
		for _, char := range row {
			if char == '#' {
				updateMinWord(current_word)
				current_word = ""
			} else {
				current_word += string(char)
			}
		}
		updateMinWord(current_word)
	}

	for j := 0; j < c; j++ {
		current_word := ""
		for i := 0; i < r; i++ {
			char := rune(grid[i][j])
			if char == '#' {
				updateMinWord(current_word)
				current_word = ""
			} else {
				current_word += string(char)
			}
		}
		updateMinWord(current_word)
	}

	fmt.Println(min_word)
}
