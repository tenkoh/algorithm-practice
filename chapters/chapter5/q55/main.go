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
			// a[i]を使わずにwを作れる時
			dp[i+1][w] = dp[i][w]

			// なんどでもa[i]を足せるので次のように更新
			// dp[i]からの遷移は　↑ で網羅済み
			if w-a[i] >= 0 {
				dp[i+1][w] = dp[i+1][w] || dp[i+1][w-a[i]]
			}
		}
	}
	fmt.Println(dp)
	return dp[N][W]
}
