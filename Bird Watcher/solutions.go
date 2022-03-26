package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
 	var sum int
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
} 

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
		var sum int
	if week == 1 { 
		for _, v := range birdsPerDay[:7] {
			sum += v
		}
	} else if week == 2 { 
		for _, v := range birdsPerDay[8:14] {
			sum += v
		}
	}
	return sum

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 { // i += 2, i = i + 2
		birdsPerDay[i]++ // i = i + 1
	}
	return birdsPerDay
}
