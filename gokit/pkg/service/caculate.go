package service

import (
	"context"
	"errors"
	"math/rand"
)

type Caculate interface {
	Sum(context.Context, int, int) (int, error)
}

type caculate struct{}

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

var (
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	ErrIntOverflow = errors.New("integer overflow")

	Server_hash string
)

func init() {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, 15)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	Server_hash = string(s)
}

func (s caculate) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, errors.New(Server_hash)
}

func NewCaculate() Caculate {
	return &caculate{}
}
