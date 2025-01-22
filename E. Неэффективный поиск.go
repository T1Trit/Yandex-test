package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, q int
	fmt.Fscanln(r, &n)
	db := make([]string, n)
	for i := range db {
		fmt.Fscanln(r, &db[i])
	}
	fmt.Fscanln(r, &q)

	queries := make([]string, q)
	for i := range queries {
		fmt.Fscanln(r, &queries[i])
	}

	batchSize := 3000           
	wordsPerQueryLimit := 23000 

	for batchStart := 0; batchStart < q; batchStart += batchSize {
		batchEnd := batchStart + batchSize
		if batchEnd > q {
			batchEnd = q
		}

		lcpTable := make(map[int]map[int]byte)

		for i := batchStart; i < batchEnd; i++ {
			query := queries[i]
			lcpTable[i] = make(map[int]byte)
			queryBytes := []byte(query)

			for j, word := range db {
				if j >= wordsPerQueryLimit {
					continue
				}
				lcp := calculateLCP(queryBytes, []byte(word))
				lcpTable[i][j] = byte(lcp)
			}
		}

		for i := batchStart; i < batchEnd; i++ {
			query := queries[i]
			ops := 0
			queryBytes := []byte(query)

			for j, word := range db {
				lcp := 0
				if j < wordsPerQueryLimit {
					if val, ok := lcpTable[i][j]; ok {
						lcp = int(val)
					} else {
						lcp = calculateLCP(queryBytes, []byte(word))
					}
				} else {
					lcp = calculateLCP(queryBytes, []byte(word))
				}

				ops += 1 + lcp
				if query == word {
					fmt.Fprintln(w, ops)
					goto nextQuery
				}
			}
			fmt.Fprintln(w, ops)
		nextQuery:
		}
	}
}

func calculateLCP(s1, s2 []byte) int {
	lcp := 0
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	for i := 0; i < minLen; i++ {
		if s1[i] == s2[i] {
			lcp++
		} else {
			break
		}
	}
	return lcp
}
