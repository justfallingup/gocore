package geom

import "testing"

func TestGeom_Distance(t *testing.T) {
	type fields struct {
		X1 float64
		Y1 float64
		X2 float64
		Y2 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "distance",
			fields: fields{
				X1: 1,
				Y1: 1,
				X2: 4,
				Y2: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Geom{
				X1: tt.fields.X1,
				Y1: tt.fields.Y1,
				X2: tt.fields.X2,
				Y2: tt.fields.Y2,
			}
			if got := g.Distance(); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeom_Valid(t *testing.T) {
	type fields struct {
		X1 float64
		Y1 float64
		X2 float64
		Y2 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				X1: 1,
				Y1: 1,
				X2: 4,
				Y2: 5,
			},
			want: true,
		},
		{
			name: "not valid",
			fields: fields{
				X1: -1,
				Y1: 1,
				X2: 4,
				Y2: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Geom{
				X1: tt.fields.X1,
				Y1: tt.fields.Y1,
				X2: tt.fields.X2,
				Y2: tt.fields.Y2,
			}
			if got := g.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
