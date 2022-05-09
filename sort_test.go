package air

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "standard use case",
			args: args{arr: []int{10, 99, 23, 91, 5, 6, 7, 98}, low: 0, high: 7},
			want: []int{5, 6, 7, 10, 23, 91, 98, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuicksortRecursive(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "standard use case",
			args: args{arr: []int{10, 99, 23, 91, 5, 6, 7, 98}, low: 0, high: 7},
			want: []int{5, 6, 7, 10, 23, 91, 98, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuicksortRecursive(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuicksortRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
