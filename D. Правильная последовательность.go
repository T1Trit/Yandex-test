package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	mod := 1000000007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	match := func(a, b byte) bool {
		if a == '?' && (b == ')' || b == ']' || b == '}') {
			return true
		}
		if b == '?' && (a == '(' || a == '[' || a == '{') {
			return true
		}
		if a == '(' && b == ')' {
			return true
		}
		if a == '[' && b == ']' {
			return true
		}
		if a == '{' && b == '}' {
			return true
		}
		if a == '?' && b == '?' {
			return true
		}
		return false
	}
    
    countMatch := func(a,b byte) int {
        if a == '?' && b == '?' { return 3 }
        if match(a,b) { return 1 }
        return 0
    }

	for length := 2; length <= n; length += 2 {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			for k := i + 1; k <= j; k+=2 {
				
                mult := countMatch(s[i],s[k])
                
                if mult > 0 {
				    if i+1 <= k-1{
                        if k+1 <= j{
                            dp[i][j] = (dp[i][j] + mult * dp[i+1][k-1] * dp[k+1][j]) % mod
                        } else {
                            dp[i][j] = (dp[i][j] + mult * dp[i+1][k-1]) % mod
                        }
                    } else {
                        if k+1 <= j{
                           dp[i][j] = (dp[i][j] + mult * dp[k+1][j]) % mod
                        } else {
                            dp[i][j] = (dp[i][j] + mult) % mod
                        }
                    }
                }
			}
		}
	}

	fmt.Println(dp[0][n-1])
}
