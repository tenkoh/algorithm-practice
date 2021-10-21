package q54

import (
	"fmt"

	"github.com/tenkoh/algorithm-practice/pkg/util"
)

var Inf int = 1 << 29

func Solution(W, k int, a []int) bool {
	// dp[i][w] と cnt[i][w]をつくる
	// cntは最小値を更新する
	// return dp[N][W] && cnt[N][W]<=k
	N := len(a)
	dp := make([][]bool, N+1)
	for r := range dp {
		dp[r] = make([]bool, W+1)
	}
	cnt := util.MakeTable(N+1, W+1, 0)

	for i := 0; i <= N; i++ {
		for w := 0; w <= W; w++ {
			// 何も選ばない時は和が0だけtrue
			if i == 0 {
				dp[i][w] = w == 0
				continue
			}
			// a[i-1]を加えない時
			dp[i][w] = dp[i-1][w]
			// a[i-1]を加える時をORで。この時だけカウンタを増加させる。
			if w-a[i-1] < 0 {
				continue
			}
			dp[i][w] = dp[i][w] || dp[i-1][w-a[i-1]]
			cnt[i][w] = cnt[i-1][w-a[i-1]] + 1
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
		fmt.Println(cnt[i])
	}
	return dp[N][W] && (cnt[N][W] <= k)
}
