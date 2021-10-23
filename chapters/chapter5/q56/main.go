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
			// a[i]を足さない時、dp[i][w]==InfならInf, それ以外なら0を入れる
			if dp[i][w] != Inf {
				dp[i+1][w] = 0
			}
			// a[i]を足す時
			// dp[i-1]からきた時とdp[i]からきた時で分ける
			// 遷移してこられるかの条件は、遷移前の加算回数がm未満のこと
			if i > 0 {
				if w-a[i-1] >= 0 {
					if dp[i][w-a[i-1]] < m[i-1] {
						util.Chmin(&dp[i+1][w], 1)
					}
				}
			}
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
