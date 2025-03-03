package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{0,2,1,5,3,4},
			},
			want: []int{0,1,2,4,5,3},
		},
		{
			name: "test2",
			args: args{
				nums: []int{5,0,1,2,3,4},
			},
			want: []int{4,5,0,1,2,3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := buildArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
