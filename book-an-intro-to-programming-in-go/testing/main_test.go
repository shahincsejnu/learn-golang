package main

import "testing"

func TestTotal(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first test",
			args: args{
				args: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "second one",
			args: args{
				args: []int{1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Total(tt.args.args...); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				args: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				args: []int{1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xor(tt.args.args...); got != tt.want {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}
