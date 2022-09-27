package main

type IDepartment interface {
	execute(data *patient) error
	setNext(IDepartment)
}
