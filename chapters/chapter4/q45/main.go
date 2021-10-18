package q45

type Counter struct {
	bound int
	count int
}

func NewCounter(N int) *Counter {
	c := new(Counter)
	c.bound = N
	c.count = 0
	return c
}

func (c *Counter) Calc(current, used int) {
	if current > c.bound {
		return
	}
	if used == 0b111 {
		c.count++
	}
	for i, v := range []int{3, 5, 7} {
		c.Calc(current*10+v, used|1<<i)
	}
}

func Solution(N int) int {
	c := NewCounter(N)
	c.Calc(0, 0b000)
	return c.count
}
