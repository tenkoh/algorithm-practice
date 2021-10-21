package q52

func Solution(W int, a []int) bool {
	// dp[N+1][W+1] bool
	// dp[i+1][w] = dp[i][w] || dp[i][w-a[i+1]]
	N := len(a)
	dp := make([][]bool, N+1)
	for r := range dp {
		dp[r] = make([]bool, W+1)
	}
	for i := 0; i <= N; i++ {
		for w := 0; w <= W; w++ {
			// 何も選ばない時は和が0だけtrue
			if i == 0 {
				dp[i][w] = w == 0
				continue
			}
			// a[i-1]を加えない時
			dp[i][w] = dp[i-1][w]
			// a[i-1]を加える時をORで
			if w >= a[i-1] {
				dp[i][w] = dp[i][w] || dp[i-1][w-a[i-1]]
			}
		}
	}
	return dp[N][W]
}
