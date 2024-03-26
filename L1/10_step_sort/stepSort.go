package stepsort

import "math"

func StepSort(arr []float32, step float32) map[int][]float32{
	groups := make(map[int][]float32)

	for _, v := range arr {
		key := int(math.Floor(float64(v / step))) * 10
		groups[key] = append(groups[key], v)
	}

	return groups

}  