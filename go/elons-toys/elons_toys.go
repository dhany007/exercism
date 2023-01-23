package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	battery := c.battery - c.batteryDrain
	if battery > 0 {
		c.battery = battery
		c.distance = c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	result := fmt.Sprintf("Driven %d meters", c.distance)
	return result
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	result := fmt.Sprintf("Battery at %d%s", c.battery, "%")
	return result
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	countDistance := float64(trackDistance) / float64(c.speed)
	countBattery := countDistance * float64(c.batteryDrain)

	currentBattery := float64(c.battery) - countBattery

	return currentBattery >= 0
}
