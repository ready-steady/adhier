package polynomial

import (
	"testing"

	"github.com/ready-steady/adapt/internal"
	"github.com/ready-steady/assert"

	grid "github.com/ready-steady/adapt/grid/equidistant"
)

func BenchmarkClosedCompute1(b *testing.B) {
	benchmarkClosedCompute(1, b)
}

func BenchmarkClosedCompute2(b *testing.B) {
	benchmarkClosedCompute(2, b)
}

func BenchmarkClosedCompute3(b *testing.B) {
	benchmarkClosedCompute(3, b)
}

func benchmarkClosedCompute(power uint, b *testing.B) {
	const (
		nd = 10
		ns = 100000
	)

	basis := NewClosed(nd, power)
	indices := generateIndices(nd, ns, grid.NewClosed(nd).Refine)
	points := generatePoints(nd, ns, indices, closedNode)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < ns; j++ {
			basis.Compute(indices[j*nd:(j+1)*nd], points[j*nd:(j+1)*nd])
		}
	}
}

func TestClosedCompute1D1P(t *testing.T) {
	basis := NewClosed(1, 1)

	compute := func(level, order uint64, point float64) float64 {
		return basis.Compute(internal.Compose([]uint64{level}, []uint64{order}), []float64{point})
	}

	points := []float64{0.0, 0.25, 0.5, 0.75, 1.0}

	cases := []struct {
		level  uint64
		order  uint64
		values []float64
	}{
		{0, 0, []float64{1.0, 1.0, 1.0, 1.0, 1.0}},
		{1, 0, []float64{1.0, 0.5, 0.0, 0.0, 0.0}},
		{1, 2, []float64{0.0, 0.0, 0.0, 0.5, 1.0}},
		{2, 1, []float64{0.0, 1.0, 0.0, 0.0, 0.0}},
		{2, 3, []float64{0.0, 0.0, 0.0, 1.0, 0.0}},
	}

	values := make([]float64, len(points))

	for i := range cases {
		for j := range values {
			values[j] = compute(cases[i].level, cases[i].order, points[j])
		}
		assert.Equal(values, cases[i].values, t)
	}
}

func TestClosedCompute1D2P(t *testing.T) {
	const (
		nd = 1
		np = 2
		nb = 4
		nn = 101
	)

	basis := NewClosed(nd, np)

	indices := internal.Compose(
		[]uint64{3, 3, 3, 3, 3, 3, 3, 3},
		[]uint64{1, 3, 5, 7, 9, 11, 13, 15},
	)

	points := make([]float64, nn)
	for i := range points {
		points[i] = float64(i) / (nn - 1)
	}

	values := make([]float64, nn)
	for i := range values {
		for j := 0; j < nb; j++ {
			values[i] += basis.Compute(indices[j:j+1], points[i:i+1])
		}
	}

	assert.EqualWithin(values, []float64{
		0.0000000000000000e+00, 1.5359999999999999e-01, 2.9440000000000000e-01,
		4.2240000000000000e-01, 5.3759999999999997e-01, 6.4000000000000012e-01,
		7.2960000000000003e-01, 8.0640000000000001e-01, 8.7039999999999995e-01,
		9.2159999999999997e-01, 9.5999999999999996e-01, 9.8560000000000014e-01,
		9.9839999999999995e-01, 9.9839999999999995e-01, 9.8559999999999992e-01,
		9.5999999999999996e-01, 9.2159999999999997e-01, 8.7039999999999995e-01,
		8.0640000000000001e-01, 7.2960000000000003e-01, 6.3999999999999990e-01,
		5.3760000000000008e-01, 4.2240000000000000e-01, 2.9439999999999988e-01,
		1.5360000000000013e-01, 0.0000000000000000e+00, 1.5360000000000013e-01,
		2.9440000000000022e-01, 4.2240000000000033e-01, 5.3759999999999974e-01,
		6.3999999999999990e-01, 7.2960000000000003e-01, 8.0640000000000001e-01,
		8.7040000000000006e-01, 9.2160000000000009e-01, 9.5999999999999996e-01,
		9.8559999999999992e-01, 9.9839999999999995e-01, 9.9839999999999995e-01,
		9.8559999999999992e-01, 9.5999999999999996e-01, 9.2160000000000009e-01,
		8.7040000000000006e-01, 8.0640000000000001e-01, 7.2960000000000003e-01,
		6.3999999999999990e-01, 5.3759999999999974e-01, 4.2240000000000033e-01,
		2.9440000000000022e-01, 1.5360000000000013e-01, 0.0000000000000000e+00,
		1.5360000000000013e-01, 2.9440000000000022e-01, 4.2240000000000033e-01,
		5.3760000000000041e-01, 6.4000000000000046e-01, 7.2960000000000047e-01,
		8.0639999999999967e-01, 8.7039999999999973e-01, 9.2159999999999986e-01,
		9.5999999999999996e-01, 9.8559999999999992e-01, 9.9839999999999995e-01,
		9.9839999999999995e-01, 9.8559999999999992e-01, 9.5999999999999996e-01,
		9.2159999999999986e-01, 8.7039999999999973e-01, 8.0639999999999967e-01,
		7.2960000000000047e-01, 6.4000000000000046e-01, 5.3760000000000041e-01,
		4.2240000000000033e-01, 2.9440000000000022e-01, 1.5360000000000013e-01,
		0.0000000000000000e+00, 1.5360000000000013e-01, 2.9440000000000022e-01,
		4.2240000000000033e-01, 5.3760000000000041e-01, 6.4000000000000046e-01,
		7.2960000000000047e-01, 8.0639999999999967e-01, 8.7039999999999973e-01,
		9.2159999999999986e-01, 9.5999999999999996e-01, 9.8559999999999992e-01,
		9.9839999999999995e-01, 9.9839999999999995e-01, 9.8559999999999992e-01,
		9.5999999999999996e-01, 9.2159999999999986e-01, 8.7039999999999973e-01,
		8.0639999999999967e-01, 7.2960000000000047e-01, 6.4000000000000046e-01,
		5.3760000000000041e-01, 4.2240000000000033e-01, 2.9440000000000022e-01,
		1.5360000000000013e-01, 0.0000000000000000e+00,
	}, 1e-15, t)
}

func TestClosedIntegrate(t *testing.T) {
	basis := NewClosed(1, 1)

	levels := []uint64{0, 1, 2, 3}
	values := []float64{1.0, 0.25, 1.0 / 2.0 / 2.0, 1.0 / 2.0 / 2.0 / 2.0}

	for i := range levels {
		indices := internal.Compose([]uint64{levels[i]}, []uint64{0})
		assert.Equal(basis.Integrate(indices), values[i], t)
	}
}

func TestClosedParent(t *testing.T) {
	childLevels := []uint64{1, 1, 2, 2, 3, 3, 3, 3}
	childOrders := []uint64{0, 2, 1, 3, 1, 3, 5, 7}

	parentLevels := []uint64{0, 0, 1, 1, 2, 2, 2, 2}
	parentOrders := []uint64{0, 0, 0, 2, 1, 1, 3, 3}

	for i := range childLevels {
		level, order := closedParent(childLevels[i], childOrders[i])
		assert.Equal(level, parentLevels[i], t)
		assert.Equal(order, parentOrders[i], t)
	}
}
