package react

import (
	"reflect"
	"runtime"
)

// MyReactor implements Reactor
type MyReactor struct{}

var pubsub = map[Cell][]*MyComputeCell{}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (m *MyReactor) CreateInput(val int) InputCell {
	input := &MyInputCell{}
	input.SetValue(val)
	return input
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (m *MyReactor) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	computeCell := &MyComputeCell{}

	wrapper := func(comp *MyComputeCell, in Cell) func(int) {
		return func(val int) {
			comp.val = compute(val)

			for _, sub := range pubsub[comp] {
				for _, cb := range sub.callbacks {
					cb(comp.val)
				}
			}

			for i, cb := range comp.callbacks {
				if i == 0 {
					continue
				}
				cb(comp.val)
			}
		}
	}

	cb := wrapper(computeCell, cell)
	cb(cell.Value())
	computeCell.AddCallback(cb)
	computeCell.inputs = append(computeCell.inputs, cell)

	associateInputCompute(cell, computeCell)

	return computeCell
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (m *MyReactor) CreateCompute2(cell1 Cell, cell2 Cell, compute func(val1 int, val2 int) int) ComputeCell {
	computeCell := &MyComputeCell{}

	wrapper := func(comp *MyComputeCell, in1 Cell, in2 Cell) func(int) {
		return func(val int) {
			comp.val = compute(in1.Value(), in2.Value())

			for _, sub := range pubsub[comp] {
				for _, cb := range sub.callbacks {
					cb(comp.val)
				}
			}
		}
	}

	cb := wrapper(computeCell, cell1, cell2)
	cb(cell1.Value() + cell2.Value())
	computeCell.AddCallback(cb)

	computeCell.inputs = append(computeCell.inputs, cell1)
	computeCell.inputs = append(computeCell.inputs, cell2)

	associateInputCompute(cell1, computeCell)
	associateInputCompute(cell2, computeCell)

	return computeCell
}

// MyCell implements Cell
type MyCell struct {
	val int
}

// Value returns the current value of the cell.
func (m *MyCell) Value() int {
	return m.val
}

// MyInputCell implements InputCell
type MyInputCell struct {
	MyCell
}

// SetValue sets the value of the cell.
func (m *MyInputCell) SetValue(val int) {
	m.val = val

	for _, sub := range pubsub[m] {
		for i, cb := range sub.callbacks {
			if i != 0 {
				continue
			}
			cb(m.Value())
		}
	}
}

// MyComputeCell implements ComputeCell
type MyComputeCell struct {
	MyCell
	inputs    []Cell
	callbacks []func(int)
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (m *MyComputeCell) AddCallback(cb func(int)) Canceler {
	m.callbacks = append(m.callbacks, cb)

	return &MyCanceler{
		cb,
	}
}

// MyCanceler implements Canceler
type MyCanceler struct {
	callback func(int)
}

// Cancel removes the callback.
func (m *MyCanceler) Cancel() {
	m.callback = nil
}

// New returns a new Reactor
func New() Reactor {
	return &MyReactor{}
}

/// helper funcs ////
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func associateInputCompute(pub Cell, subscriber *MyComputeCell) {
	pubsub[pub] = append(pubsub[pub], subscriber)
}
