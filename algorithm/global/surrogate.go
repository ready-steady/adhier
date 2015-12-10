package global

import (
	"fmt"
)

// Surrogate is an interpolant for a function.
type Surrogate struct {
	Inputs    uint      // Number of inputs
	Outputs   uint      // Number of outputs
	Nodes     uint      // Number of nodes
	Indices   []uint64  // Indices of the nodes
	Surpluses []float64 // Hierarchical surpluses
}

func newSurrogate(ni, no uint) *Surrogate {
	return &Surrogate{
		Inputs:    ni,
		Outputs:   no,
		Indices:   make([]uint64, 0),
		Surpluses: make([]float64, 0),
	}
}

func (self *Surrogate) push(indices []uint64, surpluses []float64) {
	na := uint(len(indices)) / self.Inputs
	self.Nodes += na
	self.Indices = append(self.Indices, indices...)
	self.Surpluses = append(self.Surpluses, surpluses...)
}

// String returns a human-friendly representation.
func (self *Surrogate) String() string {
	phantom := struct {
		inputs  uint
		outputs uint
		nodes   uint
	}{
		inputs:  self.Inputs,
		outputs: self.Outputs,
		nodes:   self.Nodes,
	}
	return fmt.Sprintf("%+v", phantom)
}