package main

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				jewels: "z",
				stones: "ZZ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
