package c54

import (
	"github.com/tenkoh/algorithm-practice/pkg/util"
)

var Inf int = 1 << 30

func Solution(N int, h []int) int {
	dp := make([]int, N)
	for i := range dp {
		dp[i] = Inf
	}

	dp[0] = 0
	for idx := 1; idx < N; idx++ {
		util.ChooseInt(&dp[idx], dp[idx-1]+util.AbsInt(h[idx]-h[idx-1]), "min")
		if idx < 2 {
			continue
		}
		util.ChooseInt(&dp[idx], dp[idx-2]+util.AbsInt(h[idx]-h[idx-2]), "min")
	}
	return dp[N-1]
}
