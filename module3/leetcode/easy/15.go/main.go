package main

import "fmt"

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.big > 0 {
			ps.big--
			return true
		}
	case 2:
		if ps.medium > 0 {
			ps.medium--
			return true
		}
	case 3:
		if ps.small > 0 {
			ps.small--
			return true
		}
	}
	return false
}

func main() {
	parkingSystem := Constructor(1, 1, 0)
	fmt.Println(parkingSystem.AddCar(1)) // Output: true
	fmt.Println(parkingSystem.AddCar(2)) // Output: true
	fmt.Println(parkingSystem.AddCar(3)) // Output: false
	fmt.Println(parkingSystem.AddCar(1)) // Output: false
}
