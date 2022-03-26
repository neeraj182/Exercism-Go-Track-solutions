package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 { 
    var cars_success float64 = float64(productionRate) * float64(successRate/100)
    return cars_success
	
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
   
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var cars float64 = float64 (productionRate / 60)
    var car float64 = (successRate/100) * cars
    var carss int = int (car)
	return carss  
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
     var cars_no = carsCount / 10
	var remaining_cars = carsCount % 10
	var total_cost = (uint(cars_no) * 95000) + (uint(remaining_cars) * 10000)
	return total_cost

    
	
}
