package geom

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "distance",
			args: args{
				1.0,
				1.0,
				4.0,
				5.0},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}