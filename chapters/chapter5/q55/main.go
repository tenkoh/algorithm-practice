package q55

import "fmt"

func Solution(W int, a []int) bool {
	N := len(a)
	dp := make([][]bool, N+1)
	for r := range dp {
		dp[r] = make([]bool, W+1)
	}
	dp[0][0] = true
	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// a[i]を選ばない時
			dp[i+1][w] = dp[i][w]

			// a[i]を選ぶ時、その前にa[i]を選んだ時とa[i-1]を選んだ時の二通りがある
			if w-a[i] >= 0 {
				dp[i+1][w] = dp[i+1][w] || dp[i+1][w-a[i]]
			}
			if i < 1 {
				continue
			}
			if w-a[i-1] >= 0 {
				dp[i+1][w] = dp[i+1][w] || dp[i][w-a[i-1]]
			}
		}
	}
	fmt.Println(dp)
	return dp[N][W]
}
