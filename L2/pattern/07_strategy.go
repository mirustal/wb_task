package pattern

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct {
}

func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

type Context struct {
	strategy StrategySort
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
