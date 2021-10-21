package c57

import (
	"github.com/tenkoh/algorithm-practice/pkg/util"
)

func Solution(N, W int, weight, value []int) int {
	// dp tableをつくる
	// 入れた時入れない時　×　重さ制約の組合せに対して価値の変動テーブルを作る
	// 貰う形式で推移させる
	dp := util.MakeTable(N+1, W+1, 0)
	// choose nothing
	dp[0][0] = 0
	for item := 0; item < N; item++ {
		for limit := 0; limit <= W; limit++ {
			// when choose the item
			if nw := limit - weight[item]; nw >= 0 {
				util.Chmax(&dp[item+1][limit], dp[item][nw]+value[item])
			}
			// when not choose the item
			util.Chmax(&dp[item+1][limit], dp[item][limit])
		}
	}

	return dp[N][W]
}
