package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    totalBird := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalBird += birdsPerDay[i]
	}

	return totalBird
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var (
		dayPerWeek int = 7
		totalBirds int = 0
	)

	firstDay := dayPerWeek*(week-1) + 1
	lastDay := dayPerWeek * week

	newBirdsWeek := birdsPerDay[firstDay-1 : lastDay]

	for i := 0; i < len(newBirdsWeek); i++ {
		totalBirds += newBirdsWeek[i]
	}

	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 1; i < len(birdsPerDay)+1; i++ {
		if i%2 != 0 {
			birdsPerDay[i-1] = birdsPerDay[i-1] + 1
		}
	}

	return birdsPerDay
}
