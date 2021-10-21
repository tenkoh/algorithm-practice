package q51

import "github.com/tenkoh/algorithm-practice/pkg/util"

func Solution(N int, a, b, c []int) int {
	// dp[i][j]: j=0,1,2 => a,b,c
	// dp[i][0] = max(dp[i-1][1] + a[i], dp[i-1][2] + a[i])
	// return max(dp[N-1][j])
	dp := util.MakeTable(N, 3, 0)
	dp[0] = []int{a[0], b[0], c[0]}
	for i := 1; i < N; i++ {
		util.Chmax(&dp[i][0], dp[i-1][1]+a[i])
		util.Chmax(&dp[i][0], dp[i-1][2]+a[i])
		util.Chmax(&dp[i][1], dp[i-1][2]+b[i])
		util.Chmax(&dp[i][1], dp[i-1][0]+b[i])
		util.Chmax(&dp[i][2], dp[i-1][0]+c[i])
		util.Chmax(&dp[i][2], dp[i-1][1]+c[i])
	}
	m := 0
	util.Chmax(&m, dp[N-1][0])
	util.Chmax(&m, dp[N-1][1])
	util.Chmax(&m, dp[N-1][2])
	return m
}
