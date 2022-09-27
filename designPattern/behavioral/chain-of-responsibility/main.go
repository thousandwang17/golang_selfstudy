package main

func main() {
	patient := NewPatient("test")

	hospital(patient)
}

func hospital(patient *patient) {

	reception := NewReception()

	medical := NewMedical()
	medical.setNext(reception)

	doctor := NewDoctor()
	doctor.setNext(medical)

	cashier := NewCashier()
	cashier.setNext(doctor)

	cashier.execute(patient)
}

// 適合處理多線復用
// data -> 1 ->  2  -> 5
//	  		 ->  3
//    		 ->  4  -> 5
//
