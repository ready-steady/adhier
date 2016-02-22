package local

import (
	"math"
)

// Target is a function to be interpolated.
type Target interface {
	// Dimensions returns the number of inputs and the number of outputs.
	Dimensions() (uint, uint)

	// Before gets called once per iteration before involving Compute. If the
	// function returns false, the interpolation process is terminated.
	Before(*Progress) bool

	// Compute evaluates the target function at a point. The function is called
	// multiple times per iteration, depending on the number of active nodes.
	Compute(point, value []float64)

	// Score assigns a score to a location. The function is called after
	// Compute, and it is called as many times as Compute. If the score is
	// positive, the node is refined; otherwise, no refinement is performed.
	Score(*Location) float64

	// After gets called once per iteration after involving Compute and Score.
	// If the function returns false, the interpolation process is terminated.
	After(*Progress) bool
}

// Location contains information about a spacial location.
type Location struct {
	Value   []float64 // Target-function value
	Surplus []float64 // Hierarchical surplus
	Volume  float64   // Volume under the basis function
}

// Progress contains information about the interpolation process.
type Progress struct {
	Level    uint      // Reached level
	Active   uint      // Number of active nodes
	Passive  uint      // Number of passive nodes
	Refined  uint      // Number of refined nodes
	Integral []float64 // Integral over the whole domain
}

// BasicTarget is a basic target satisfying the Target interface.
type BasicTarget struct {
	Inputs  uint // > 0
	Outputs uint // > 0

	Tolerance float64 // ≥ 0

	BeforeHandler  func(*Progress) bool
	ComputeHandler func([]float64, []float64) // != nil
	ScoreHandler   func(*Location) float64
	AfterHandler   func(*Progress) bool
}

// NewTarget creates a basic target.
func NewTarget(inputs, outputs uint, tolerance float64,
	compute func([]float64, []float64)) *BasicTarget {

	return &BasicTarget{
		Inputs:  inputs,
		Outputs: outputs,

		Tolerance: tolerance,

		ComputeHandler: compute,
	}
}

func (self *BasicTarget) Dimensions() (uint, uint) {
	return self.Inputs, self.Outputs
}

func (self *BasicTarget) Before(progress *Progress) bool {
	if self.BeforeHandler != nil {
		return self.BeforeHandler(progress)
	} else {
		return true
	}
}

func (self *BasicTarget) Compute(node, value []float64) {
	self.ComputeHandler(node, value)
}

func (self *BasicTarget) Score(location *Location) float64 {
	if self.ScoreHandler != nil {
		return self.ScoreHandler(location)
	} else {
		return self.defaultScore(location)
	}
}

func (self *BasicTarget) After(progress *Progress) bool {
	if self.AfterHandler != nil {
		return self.AfterHandler(progress)
	} else {
		return true
	}
}

func (self *BasicTarget) defaultScore(location *Location) float64 {
	for _, ε := range location.Surplus {
		if math.Abs(ε) > self.Tolerance {
			return 1.0
		}
	}
	return 0.0
}
