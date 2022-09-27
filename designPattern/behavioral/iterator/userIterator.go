package main

type UserIterator struct {
	index int
	users *UserCollection
}

func NewUserIterator(u *UserCollection) iterator {
	return &UserIterator{
		users: u,
	}
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users.user)
}

func (u *UserIterator) getNext() *user {
	if u.HasNext() {
		res := u.users.user[u.index]
		u.index++
		return res
	}
	return nil
}
