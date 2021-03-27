package react

// New creates a new reactor
func New() Reactor {
	return &reactor{}
}

type reactor struct {
	computeCells []*computeCell
}

func (r *reactor) CreateInput(val int) InputCell {
	cell := &inputCell{}
	cell.reactor = r
	cell.SetValue((val))
	return cell

}

func (r *reactor) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	computeCell := r.newComputeCell(func() int {
		return compute(cell.Value())
	})
	return computeCell
}

func (r *reactor) CreateCompute2(cell1 Cell, cell2 Cell, compute func(val1 int, val2 int) int) ComputeCell {
	computeCell := r.newComputeCell(func() int {
		return compute(cell1.Value(), cell2.Value())
	})

	return computeCell
}

// helper func to make a compute cell
func (r *reactor) newComputeCell(compute func() int) *computeCell {
	computeCell := &computeCell{}
	computeCell.compute = compute
	computeCell.callbacks = make(map[int]func(int))
	computeCell.prevVal = computeCell.Value()
	r.computeCells = append(r.computeCells, computeCell)
	return computeCell
}

type inputCell struct {
	val     int
	reactor *reactor
}

func (ic *inputCell) Value() int {
	return ic.val
}

func (ic *inputCell) SetValue(val int) {
	ic.val = val

	for _, computeCell := range ic.reactor.computeCells {
		if len(computeCell.callbacks) == 0 || computeCell.Value() == computeCell.prevVal {
			continue
		}
		for _, cb := range computeCell.callbacks {
			cb(computeCell.Value())
		}
		computeCell.prevVal = computeCell.Value()
	}
}

type computeCell struct {
	prevVal   int
	compute   func() int
	callbacks map[int]func(int)
}

func (cc *computeCell) Value() int {
	return cc.compute()
}

func (cc *computeCell) AddCallback(cb func(int)) Canceler {
	cc.callbacks[len(cc.callbacks)] = cb

	return &canceler{
		computeCell: cc,
		key:         len(cc.callbacks) - 1,
	}
}

type canceler struct {
	computeCell *computeCell
	key         int
}

func (c *canceler) Cancel() {
	delete(c.computeCell.callbacks, c.key)
}
