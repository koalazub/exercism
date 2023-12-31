package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.

func (c *canceler) Cancel() {
	return c
}

func (c *cell) Value() int {
	return c
}

func (c *cell) SetValue(value int) {
	return c.InputCell(value)
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	panic("Please implement the AddCallback function")
}

func New() Reactor {
	panic("Please implement the New function")
}

func (r *reactor) CreateInput(initial int) InputCell {
	panic("Please implement the CreateInput function")
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	panic("Please implement the CreateCompute1 function")
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	panic("Please implement the CreateCompute2 function")
}
