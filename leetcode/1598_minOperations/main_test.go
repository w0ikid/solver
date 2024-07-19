package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				logs: []string{"d1/", "../", "../", "../"},
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				logs: []string{"./", "../", "./"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := minOperations(tt.args.logs); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
