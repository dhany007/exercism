// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	street, ok := r.Address["street"]
	if !ok {
		return false
	}

	if street != "" && r.Name != "" {
		return true
	}

	return false
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var result int = 0

	for _, r := range residents {
		_, ok := r.Address["street"]
		if r.Name != "" && r.Age != 0 && ok {
			result += 1
		}
	}

	return result
}
