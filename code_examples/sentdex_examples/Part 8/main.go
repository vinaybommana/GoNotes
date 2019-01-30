package main

import "fmt"

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8000", nil)
// }

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 // min 0 max 65535
	brakePedal    uint16
	steeringWheel int16 // -32k --> +32k
	topSpeedKmh   float64
}

// we don't need to change the values
// in `myCar` struct
// value receiver
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

// change the topSpeed of the vehicle
// Pointer receiver
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

// this serves the same purpose
// but not good
func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	myCar := car{
		gasPedal:      22341,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0,
	}

	fmt.Println(myCar.gasPedal)
	fmt.Println(myCar.kmh())
	fmt.Println(myCar.mph())
	myCar.newTopSpeed(500)
	fmt.Println(myCar.kmh())
	fmt.Println(myCar.mph())
}
