package binarysearch

import ("testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		search int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search 0", args: args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, search: 0}, want: 0},
		{name: "search 3", args: args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, search: 3}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, 0, len(tt.args.nums)-1, tt.args.search); got != tt.want {
				t.Errorf("orig array: %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}