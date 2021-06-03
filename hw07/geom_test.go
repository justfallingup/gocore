package geom

import (
	"reflect"
	"testing"
)

func TestGeom_Distance(t *testing.T) {
	type fields struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "distance",
			fields: fields{
				x1: 1,
				y1: 1,
				x2: 4,
				y2: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Geom{
				x1: tt.fields.x1,
				y1: tt.fields.y1,
				x2: tt.fields.x2,
				y2: tt.fields.y2,
			}
			if got := g.Distance(); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Geom
		wantErr bool
	}{
		{
			name: "new positive coordinates",
			args: args{
				x1: 1,
				y1: 1,
				x2: 4,
				y2: 5,
			},
			want: &Geom{
				x1: 1,
				y1: 1,
				x2: 4,
				y2: 5,
			},
			wantErr: false,
		},
		{
			name: "new negative coordinates",
			args: args{
				x1: -1,
				y1: 1,
				x2: 4,
				y2: 5,
			},
			want: &Geom{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

