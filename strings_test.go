package air

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		txt string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard use case",
			args: args{txt: "Hello again, my old friend.", sep: " "},
			want: []string{"Hello", "again,", "my", "old", "friend."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.txt, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args struct {
		txt []string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal use case",
			args: args{txt: []string{"Hello", "World!"}, sep: " "},
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.txt, tt.args.sep); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoDuplicatedCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard use case",
			args: args{s: "Heello     Woooooorld!"},
			want: "Helo World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoDuplicatedCharacter(tt.args.s); got != tt.want {
				t.Errorf("NoDuplicatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
