package square

import (
	"sync"
)

func Square(nums []int)  []int {
	wg := new(sync.WaitGroup)
	result := make([]int, len(nums))
	wg.Add(len(nums))
	for i, num := range nums {
		go func(i, num int) {
			defer wg.Done()
			result[i] = num * num
		}(i, num)
	}
	wg.Wait()

	return result
}