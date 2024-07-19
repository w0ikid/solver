package main

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				nums: []int{1,2,3,1,1,3},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1,1,1,1},
			},
			want: 6,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1,2,3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := numIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
