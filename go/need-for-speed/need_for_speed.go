package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	tempBattery := car.battery - car.batteryDrain
	tempDistance := car.distance + car.speed

	if tempBattery < 0 {
		tempBattery = car.battery
		tempDistance = car.distance
	}

	return Car{
		battery:      tempBattery,
		distance:     tempDistance,
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	result := false

	distanceCar := (car.battery / car.batteryDrain) * car.speed

	if distanceCar >= track.distance {
		result = true
	}

	return result
}
