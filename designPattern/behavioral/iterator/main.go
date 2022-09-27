package main

import "fmt"

func main() {
	user1 := &user{id: 1, name: "denny"}
	user2 := &user{id: 2, name: "vivi"}

	u := NewUserCollection()
	u.Insert(user1, user2)

	iterator := u.CreateIterator(Normal)

	for iterator.HasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}

	user3 := &user{id: 3, name: "denny2"}
	user4 := &user{id: 4, name: "vivi2"}
	u.Insert(user3, user4)

	for iterator.HasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
