package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
    if kind == "car" || kind == "truck" {
        return true
    }

    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
    var vehicle string

	if option1 < option2 {
		vehicle = option1
	} else {
		vehicle = option2
	}

	result := fmt.Sprintf("%s is clearly the better choice.", vehicle)
	return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64

	if age <= 3 {
		price = 80 * originalPrice / 100
	} else if age > 3 && age < 10 {
		price = 70 * originalPrice / 100
	} else {
		price = 50 * originalPrice / 100
	}

    return price
}
