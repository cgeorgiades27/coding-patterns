package largestcontainer

import "testing"

func TestLargestcontainer(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums: []int{3, 3, 3, 3},
			},
			want: 9,
		},
		{
			name: "5",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestcontainer(tt.args.nums); got != tt.want {
				t.Errorf("largestcontainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
