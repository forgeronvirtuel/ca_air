package air

import "testing"

func TestIsAllDuplicated(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name:  "standard use case",
			args:  args{numbers: []int{2, 3, 1, 4, 3, 2, 1, 4, 5, 6, 5, 6}},
			want:  true,
			want1: 0,
		},
		{
			name:  "standard use case",
			args:  args{numbers: []int{2, 3, 1, 3, 2, 1, 4, 5, 6, 5, 6}},
			want:  false,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsAllDuplicated(tt.args.numbers)
			if got != tt.want {
				t.Errorf("IsAllDuplicated() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsAllDuplicated() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
