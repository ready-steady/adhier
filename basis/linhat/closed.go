package linhat

// Closed represents an instance of the basis in [0, 1]^n.
type Closed struct {
	nd int
}

// NewClosed creates an instance of the basis in [0, 1]^n.
func NewClosed(dimensions uint) *Closed {
	return &Closed{int(dimensions)}
}

// Compute evaluates the value of a basis function at a point.
func (c *Closed) Compute(index []uint64, point []float64) float64 {
	nd := c.nd

	value := 1.0

	for i := 0; i < nd; i++ {
		level := 0xFFFFFFFF & index[i]
		if level == 0 {
			continue // value *= 1
		}

		order := index[i] >> 32

		scale := float64(uint64(2) << (level - 1))
		distance := point[i] - float64(order)/scale
		if distance < 0 {
			distance = -distance
		}
		if scale*distance < 1 {
			value *= 1 - scale*distance
		} else {
			return 0 // value *= 0
		}
	}

	return value
}

// Integrate computes the integral of a basis function in [0, 1]^n.
func (c *Closed) Integrate(index []uint64) float64 {
	nd := c.nd

	value := 1.0

	for i := 0; i < nd; i++ {
		level := 0xFFFFFFFF & index[i]
		switch level {
		case 0:
			// value *= 1
		case 1:
			value *= 0.25
		default:
			value *= 1 / float64(uint64(2)<<(level-1))
		}
	}

	return value
}
