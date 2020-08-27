package gobasis

import "fmt"

// BSplineBasis object
type BSplineBasis struct {
	order int
	knts  []float64
}

// Create BSplineBasis object
func Create(knots []float64, order int) (*BSplineBasis, error) {

	var prev float64 = knots[0]

	// check minimum size
	if len(knots) < 2*order {
		return nil, fmt.Errorf("knot vector size is %d but needs to contain 2*order sorted elements", len(knots))
	}

	// verify correct knot vector
	for _, k := range knots {
		if k < prev {
			return nil, fmt.Errorf("knot vector need to be in an non decreasing order")
		}
		prev = k
	}

	return &BSplineBasis{order: order, knts: knots}, nil
}

// Interval of the BSplineBasis function
func (bs *BSplineBasis) Interval() (float64, float64) {
	return bs.knts[0], bs.knts[len(bs.knts)-1]
}

// Eval Evaluate the BSpline basis funtcions for a given t
// returns index of first basis element and
func (bs *BSplineBasis) Eval(t float64) (int, []float64) {

	coefs := make([]float64, bs.order)
	basis := make([]float64, bs.order)
	i := bs.order - 1

	stop := len(bs.knts) - bs.order

	// locate correct starting knot index
	for k := bs.order; k < stop; k++ {
		if t == bs.knts[k] {
			i = k
			break
		}
		if t > bs.knts[k] {
			i = k
		}
	}

	// cover case if inside or outside interval
	if (t >= bs.knts[0]) && (t <= bs.knts[len(bs.knts)-1]) {
		basis[0] = 1.0
	} else {
		basis[0] = 0.0
	}

	// calculate basis function
	for j := 1; j < bs.order; j++ {
		// calculate weight coefs
		for r := 0; r <= j; r++ {
			coefs[j-r] = (t - bs.knts[i-r]) / (bs.knts[i-r+j] - bs.knts[i-r])
		}

		// calculate first term, spesical case
		basis[j] = coefs[j] * basis[j-1]

		// calculate both terms
		for r := j - 1; r > 0; r-- {
			basis[r] = coefs[r]*basis[r-1] + (1.0-coefs[r+1])*basis[r]
		}

		// calculate last term, special case
		basis[0] = (1.0 - coefs[1]) * basis[0]
	}

	i -= (bs.order - 1)

	return i, basis
}
