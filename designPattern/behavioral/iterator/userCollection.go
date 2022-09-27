package main

type UserCollection struct {
	user []*user
}

const (
	Normal = 1
)

func NewUserCollection() collection {
	return &UserCollection{}
}

func (u *UserCollection) Insert(i ...*user) {
	u.user = append(u.user, i...)
}

func (u *UserCollection) CreateIterator(t int) iterator {
	var i iterator
	switch t {
	case Normal:
		i = NewUserIterator(u)
	default:
		i = NewUserIterator(u)
	}

	return i
}
