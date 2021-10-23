package q56

import (
	"fmt"

	"github.com/tenkoh/algorithm-practice/pkg/util"
)

var Inf int = 1000

func Solution(W int, a, m []int) bool {
	N := len(a)
	// a[i]の最小加算回数を取得していく。ただし使える回数を上回ったらInfにする。
	dp := util.MakeTable(N+1, W+1, Inf)
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// a[i]を足さずにwが作れる時==dp[i][w]<Infならdp[i+1][w]=0
			if dp[i][w] < Inf {
				dp[i+1][w] = 0
			}
			// dp[i+1][w] < m[i]ならa[i]を加算できる
			// 到達不可能な組み合わせ=Infなので、同じ処理でスキップできる
			if w-a[i] >= 0 {
				if pre := dp[i+1][w-a[i]]; pre < m[i] {
					util.Chmin(&dp[i+1][w], pre+1)
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[N][W] != Inf
}
