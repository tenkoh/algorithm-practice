package q46

type Planner map[Pair]bool

func NewPlanner() Planner {
	return Planner{}
}

// [a0, a1, a2]: Pair{3, xx}
// []: Pair{0, xx}
type Pair struct {
	includeNum int
	sum        int
}

func NewPair(includeNum, sum int) Pair {
	return Pair{includeNum, sum}
}

func (p Planner) PartialSum(b []int, want int) {
	n := len(b)
	pair := NewPair(n, want)
	if n == 0 {
		p[pair] = want == 0
		return
	}
	if _, exist := p[pair]; exist {
		return
	}

	inc := NewPair(n-1, want-b[n-1])
	exc := NewPair(n-1, want)
	for _, pp := range []Pair{inc, exc} {
		if _, exist := p[pp]; !exist {
			p.PartialSum(b[:n-1], pp.sum)
		}
		if p[pp] {
			p[pair] = true
			return
		}
	}
	p[pair] = false
}

func Solution(a []int, W int) bool {
	planner := NewPlanner()
	planner.PartialSum(a, W)

	for _, v := range planner {
		if v {
			return true
		}
	}
	return false
}
