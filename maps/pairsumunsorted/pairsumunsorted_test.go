package pairsumunsorted

import (
	"reflect"
	"testing"
)

func Test_pairsumunsorted(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{-1, 3, 4, 2},
				target: 3,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairsumunsorted(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairsumunsorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
