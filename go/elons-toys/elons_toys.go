package elon

import "fmt"

// Implement the `Drive` method on the `Car` that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:
func (c *Car) Drive() {
	if c.battery <= 0 || c.battery < c.batteryDrain {
		c.distance = 0
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
// "Driven 0 meters"
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
// really need to focus on math modeling
func (c *Car) CanFinish(trackDistance int) bool {
	drivenDistance := (c.battery * c.speed) / c.batteryDrain
	return trackDistance <= drivenDistance

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
