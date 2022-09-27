package main

func main() {
	shirtItem := NewItem("Nike")

	denny := NewCustomer(1)
	vivi := NewCustomer(2)

	shirtItem.Register(denny)
	shirtItem.Register(vivi)
	denny.setId(5)

	shirtItem.UpdateAvailability()
}
