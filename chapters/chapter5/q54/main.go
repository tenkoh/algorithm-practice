package q54

import (
	"github.com/tenkoh/algorithm-practice/pkg/util"
)

var Inf int = 1 << 29

func Solution(W, k int, a []int) bool {
	// dp[i][w] と cnt[i][w]をつくる
	// cntは最小値を更新する
	// return dp[N][W] && cnt[N][W]<=k
	N := len(a)
	dp := util.MakeTable(N+1, W+1, Inf)
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// a[i]を加えない時
			util.Chmin(&dp[i+1][w], dp[i][w])
			// a[i]を加える時カウンタを増加させる。
			if w-a[i] < 0 {
				continue
			}
			util.Chmin(&dp[i+1][w], dp[i][w-a[i]]+1)
		}
	}
	return dp[N][W] <= k
}
