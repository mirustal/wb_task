package quicksort

import (
	"reflect"
	"testing"
)


func TestQuickSort(t *testing.T) {
	tests := []struct {
		args   []int
		want []int
	}{
		{[]int{5, 4, 3, 8, 6, 0, 1, 2, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{7, 8, 6, 5, 3, 4, 1, 2, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		QuickSort(tt.args, 0, len(tt.args)-1)
		got := tt.args
		if !reflect.DeepEqual(tt.args, tt.want) {
			t.Errorf("Arr: %v, want: %v", got, tt.want)
		}
	}
}



