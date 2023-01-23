package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
    productionPct := float64(productionRate) * successRate / 100
    return productionPct
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
    productionPct := float64(productionRate) * successRate / 100 / 60
    return int(productionPct)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
    costOneCar := 10000
    costTenCars := 95000

	modulusGroupCars := carsCount % 10
	groupCars := (carsCount - modulusGroupCars) / 10

	productionCost := (groupCars * costTenCars) + (modulusGroupCars * costOneCar)
	return uint(productionCost)
}
