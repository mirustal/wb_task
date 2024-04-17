package sumsquare



import (
	"sync"
	"sync/atomic"
)

func SumSquare(nums []int)  uint32{
	wg := new(sync.WaitGroup)
	var res uint32
	wg.Add(len(nums))
	for i, num := range nums {
		go func(i, num int) {
			defer wg.Done()
			atomic.AddUint32(&res, uint32(num*num))
		}(i, num)
	}
	wg.Wait()

	return res
}