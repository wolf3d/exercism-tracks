package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirdCount int = 0
    for i := 0; i < len(birdsPerDay); i++ {
        totalBirdCount += birdsPerDay[i]
    }
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    const daysPerWeek int = 7    
    var firstDay int = daysPerWeek * (week - 1)
    var lastDay int = firstDay + daysPerWeek    
    birdsPerDayInWeek := birdsPerDay[firstDay:lastDay]
	return TotalBirdCount(birdsPerDayInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
        if (i + 1) % 2 != 0 {
            birdsPerDay[i] += 1
        }
    }
	return birdsPerDay
}
