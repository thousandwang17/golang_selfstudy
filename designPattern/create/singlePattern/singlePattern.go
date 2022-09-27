package main

import (
	"fmt"
	"sync"
)

type single struct{}

var (
	singleInstance *single
	locker         sync.Mutex
	once           sync.Once
)

// method 1 : use sync.Mutex
func getInstance() *single {

	if singleInstance == nil {
		locker.Lock()
		defer locker.Unlock()
		if singleInstance == nil {
			fmt.Println("created single !!")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created -2")
		}
	} else {
		fmt.Println("single instance already created -1")
	}

	return singleInstance
}

// method 2 : use sync.once

func getInstance2() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("created single !!")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("single instance already created -1")
	}

	return singleInstance
}
