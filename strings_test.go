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

