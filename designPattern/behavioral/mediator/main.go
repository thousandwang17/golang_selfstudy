package main

func main() {
	station := NewstationMannger()

	train1 := NewPassgerTrain(station)
	train2 := NewFreightTrain(station)

	train1.arrived()
	train2.arrived()
	train1.depart()
}
