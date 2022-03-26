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
	var NewCar Car
	NewCar = Car{100, batteryDrain, speed, 0}
	return NewCar

}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}
// NewTrack created a new track
func NewTrack(distance int) Track {
	var NewTrack Track
	NewTrack = Track{distance}
	return NewTrack
}
// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		car = Car{car.battery, car.batteryDrain, car.speed, car.distance}

	} else if car.batteryDrain < car.battery {
		car.distance = car.speed + car.distance
		car.battery = 100 - car.batteryDrain
		car = Car{car.battery, car.batteryDrain, car.speed, car.distance}
	}

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var TotalDistance int = (car.battery / car.batteryDrain) * car.speed
	if TotalDistance >= track.distance {
		return true
	}
	return false
}
