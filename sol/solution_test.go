package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{3, 1, 5, 8}
	for idx := 0; idx < b.N; idx++ {
		maxCoins(nums)
	}
}
func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,1,5,8]",
			args: args{nums: []int{3, 1, 5, 8}},
			want: 167,
		},
		{
			name: "nums = [1, 5]",
			args: args{nums: []int{1, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
