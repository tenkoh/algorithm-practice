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

	return 0
}
