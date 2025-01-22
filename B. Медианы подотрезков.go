package main

import "fmt"

func main() {
	var n, b int
	fmt.Scan(&n, &b)
	a := make([]int, n)
	pos := -1
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] == b {
			pos = i
		}
	}

	counts := make(map[int]int)
	balance := 0
	for i := pos; i >= 0; i-- {
		if a[i] < b {
			balance--
		} else if a[i] > b {
			balance++
		}
		counts[balance]++
	}

	ans := 0
	balance = 0
	for i := pos; i < n; i++ {
		if a[i] < b {
			balance--
		} else if a[i] > b {
			balance++
		}
		ans += counts[-balance]
	}

	fmt.Println(ans)
}
