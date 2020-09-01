package gobasis

import (
	"reflect"
	"testing"
)

func TestFirstOrder(t *testing.T) {
	coefs := []float64{0.0, 1.0}

	basis, err := Create(coefs, 1)

	if err != nil {
		t.Errorf("Initializing error")
	}

	tt := []float64{0.0, 0.5, 1.0, -0.0000001, 1.0000001}
	rr := []float64{1.0, 1.0, 1.0, 0.0, 0.0}

	for idx := 0; idx < len(tt); idx++ {
		i, result := basis.Eval(tt[idx])

		if len(result) != 1 {
			t.Errorf("expected result size: '%v' got '%v'", 1, len(result))
		}

		if i != 0 {
			t.Errorf("expected index i: '%v' got '%v'", 0, i)
		}

		if result[0] != rr[idx] {
			t.Errorf("expected result[%v]: '%v' got '%v'", i, rr[idx], result[0])
		}
	}
}

func TestDerivate(t *testing.T) {
	coefs := []float64{0.0, 0.0, 1.0, 1.0}

	basis, _ := Create(coefs, 2)

	basis2 := basis.Derivate()

	if basis2 != nil {
		if basis2.order != 1 {
			t.Errorf("expected: %v got %v", basis.order-1, basis2.order)
		}
	} else {
		t.Errorf("derivate is %v", basis2)
	}
}

func TestOrder(t *testing.T) {
	coefs := []float64{0.0, 0.0, 1.0, 1.0}

	basis, _ := Create(coefs, 2)

	order := basis.Order()

	if order != 2 {
		t.Errorf("expected: %v got %v", 2, basis.order)
	}
}

func TestBSplineBasis_Eval(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name  string
		bs    *BSplineBasis
		args  args
		want  int
		want1 []float64
	}{
		{
			"First order for interval [0, 1] when t is 0.0",
			&BSplineBasis{order: 1, knts: []float64{0.0, 1.0}}, args{t: 0.0}, 0, []float64{1.0},
		},
		{
			"First order for interval [0, 1] when t is 0.5",
			&BSplineBasis{order: 1, knts: []float64{0.0, 1.0}}, args{t: 0.5}, 0, []float64{1.0},
		},
		{
			"First order for interval [0, 1] when t is 1.0",
			&BSplineBasis{order: 1, knts: []float64{0.0, 1.0}}, args{t: 1.0}, 0, []float64{1.0},
		},
		{
			"First order for interval [0, 1] when t is utside interval from left side",
			&BSplineBasis{order: 1, knts: []float64{0.0, 1.0}}, args{t: -0.000001}, 0, []float64{0.0},
		},
		{
			"First order for interval [0, 1] when t is utside interval from right side",
			&BSplineBasis{order: 1, knts: []float64{0.0, 1.0}}, args{t: 1.000001}, 0, []float64{0.0},
		},
		{
			"hahahah",
			&BSplineBasis{order: 4, knts: []float64{0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0}},
			args{t: 0.0}, 0, []float64{1.0, 0.0, 0.0, 0.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.bs.Eval(tt.args.t)
			if got != tt.want {
				t.Errorf("BSplineBasis.Eval() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BSplineBasis.Eval() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBSplineBasis_Interval(t *testing.T) {

	basis, _ := Create([]float64{0.0, 1.0}, 1)

	tests := []struct {
		name  string
		bs    *BSplineBasis
		want  float64
		want1 float64
	}{
		{"Check Interval is reported correct", basis, 0.0, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.bs.Interval()
			if got != tt.want {
				t.Errorf("BSplineBasis.Interval() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BSplineBasis.Interval() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
