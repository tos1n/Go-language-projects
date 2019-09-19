package main

import (
	"errors"
	"fmt"
)

type Trip struct {
	destination string
	weight      float32
	deadline    int
}

type Truck struct {
	vehicle     string
	name        string
	destination string
	speed       float32
	capacity    float32
	load        float32
}
type Pickup struct {
	Truck
	isPrivate bool
}
type TrainCar struct {
	Truck
	railway string
}

func NewTruck(t *Truck) (truck Truck) {
	t.vehicle = "Truck"
	t.name = "Truck"
	t.destination = ""
	t.speed = 40
	t.capacity = 10
	t.load = 0
	return truck

}
func NewPickUp(t *Pickup) (pickup Pickup) {
	t.vehicle = "Pickup"
	t.name = "Pickup"
	t.destination = ""
	t.speed = 60
	t.capacity = 2
	t.load = 0
	t.isPrivate = true
	return pickup
}
func NewTrainCar(t *TrainCar) (traincar TrainCar) {
	t.vehicle = "TrainCar"
	t.name = "TrainCar"
	t.destination = ""
	t.speed = 30
	t.capacity = 30
	t.load = 0
	t.railway = "CNR"
	return traincar

}

type Transporter interface {
	addLoad(Trip) error
	print()
}

func (v *Truck) print() {
	fmt.Println(v.vehicle, v.name, "to", v.destination, "with", v.load, ".000000 tons")

}
func (v *Pickup) print() {
	fmt.Println(v.vehicle, v.name, "to", v.destination, "with", v.load, ".000000 tons", "(Private: ", v.isPrivate, ")")
}
func (v *TrainCar) print() {
	fmt.Println(v.vehicle, v.name, "to", v.destination, "with", v.load, ".000000 tons", "(", v.railway, ")")
}

func (v *Truck) addLoad(t Trip) error { // v stands for vehicle , t stands for trip
	var distance int
	var checkTime int

	v.load += t.weight
	if v.destination == "" {
		v.destination = t.destination
			
	} else if v.destination != t.destination {		
		v.load -= t.weight		
		return errors.New("Error: Other destination") 
	}

	if v.capacity < (v.load) {
		
		v.load -= t.weight
		
		return errors.New("Error: out of capacity") 
	}

	
	if t.destination == "Toronto" {
		distance = 400
		checkTime = distance / int(v.speed)
	}
	if t.destination == "Montreal" {
		distance = 200
		checkTime = distance / int(v.speed)
	}

	if checkTime > t.deadline { // automatically assume everything is onTime otherwise this runs
		v.load -= t.weight // removing the load that was added
		
		return errors.New("Error: There is not enough time to complete the trip") // custom string error
	}
	return nil // this return statement  signifies success

}
func (v *Pickup) addLoad(t Trip) error { // v stands for vehicle , t stands for trip
	var distance int
	var checkTime int

	v.load += t.weight
	if v.destination == "" {
			
	v.destination = t.destination
	
	} else if v.destination != t.destination {

		v.load -= t.weight
		return errors.New("Error: Other destination") 

	}
	if v.capacity < (v.load) {
		
		v.load -= t.weight
		
		return errors.New("Error: out of capacity") 
	}

	
	if t.destination == "Toronto" {
		distance = 400
		checkTime = distance / int(v.speed)
	}
	if t.destination == "Montreal" {
		distance = 200
		checkTime = distance / int(v.speed)
	}

	if checkTime > t.deadline { // automatically assume everything is onTime otherwise this runs
		v.load -= t.weight // removing the load that was added
		
		return errors.New("Error: There is not enough time to complete the trip")
	}
	return nil // this return statement  signifies success

}
func (v *TrainCar) addLoad(t Trip) error { // v stands for vehicle , t stands for trip
	var distance int
	var checkTime int

	v.load += t.weight

	if v.destination != t.destination {

		if v.destination == "" {
			v.destination = t.destination
			
		} else {
			
			v.load -= t.weight
			
			return errors.New("Error: Other destination") 

		}
	}

	if v.capacity < (v.load) {
		
		v.load -= t.weight
		
		return errors.New("Error: out of capacity") 
	}

	if t.destination == "Toronto" {
		distance = 400
		checkTime = distance / int(v.speed)
	}
	if t.destination == "Montreal" {
		distance = 200
		checkTime = distance / int(v.speed)
	}

	if checkTime > t.deadline { // automatically assume everything is onTime otherwise this runs
		v.load -= t.weight // removing the load that was added
		
		return errors.New("Error: There is not enough time to complete the trip") 
	}
	return nil // this return statement  signifies success

}

func NewTorontoTrip(weight float32, deadline int) (t *Trip) {
	a := Trip{"Toronto", weight, deadline}

	x := &a 
	
	return x
}
func NewMontrealTrip(weight float32, deadline int) (t *Trip) {
	a := Trip{"Montreal", weight, deadline}

	x := &a
	
	return x
}

func main() {
	var listOfTrips []Trip
	
	vehicle1 := Truck{}
	vehicle2 := Truck{}
	NewTruck(&vehicle1)
	NewTruck(&vehicle2)
	vehicle3 := Pickup{}
	vehicle4 := Pickup{}
	vehicle5 := Pickup{}
	NewPickUp(&vehicle3)
	NewPickUp(&vehicle4)
	NewPickUp(&vehicle5)
	vehicle6 := TrainCar{}
	NewTrainCar(&vehicle6)

	var location string
	var input_weight float32
	var input_deadline int
	
	listOfVehicles := []Transporter {&vehicle1,&vehicle2,&vehicle3,&vehicle4,&vehicle5,&vehicle6}
	

	
	for {
		fmt.Print("Destination: (t) oronto, (m) ontreal, else exit?")
		fmt.Scan(&location)
		if location == "t" || location == "Tor" || location == "tor" || location == "T" || location == "to" || location == "To" { // redant code
			fmt.Print("Weight:")
			fmt.Scan(&input_weight)
			fmt.Print("Deadline (in hours):")
			fmt.Scan(&input_deadline)
			tripA := NewTorontoTrip(input_weight, input_deadline)

			for _, i := range listOfVehicles {
				

		t1 := i.addLoad(Trip{tripA.destination, tripA.weight, tripA.deadline})
		if t1 != nil { // if the trip is unsucessfull
		
				fmt.Println(t1)


		}else{
			listOfTrips = append(listOfTrips, Trip{tripA.destination, tripA.weight, tripA.deadline})
			break
		}
         		
       
		}	
		} else if location == "m" || location == "Mon" || location == "mon" || location == "M" || location == "mo" || location == "Mo" { // redundant code
			fmt.Print("Weight:")
			fmt.Scan(&input_weight)
			fmt.Print("Deadline (in hours):")
			fmt.Scan(&input_deadline)
			tripA := NewMontrealTrip(input_weight, input_deadline)

			for _, i := range listOfVehicles {
		t1 := i.addLoad(Trip{tripA.destination, tripA.weight, tripA.deadline})
		
		if t1 != nil { // if the trip is unsucessfull
			
				fmt.Println(t1) //return approppiate error


		}else{ // if the trip is succesful
			
			listOfTrips = append(listOfTrips, Trip{tripA.destination, tripA.weight, tripA.deadline})
			break // exit loop and restart
		}

			
		}

		} else {
			fmt.Println("Not going to TO or Montreal, bye !")
			break
		}
		
	}
	fmt.Println(listOfTrips)
	fmt.Println("Vehicles:")
	vehicle1.name = "A"
	vehicle2.name = "B"
	vehicle3.name = "A"
	vehicle4.name = "B"
	vehicle5.name = "C"
	vehicle6.name = "A"
	vehicle1.print()
	vehicle2.print()
	vehicle3.print()
	vehicle4.print()
	vehicle5.print()
	vehicle6.print()
}
