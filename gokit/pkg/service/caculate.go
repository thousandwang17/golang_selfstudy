package service

import (
	"context"
	"errors"
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
)

func (s caculate) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

func NewCaculate() Caculate {
	return &caculate{}
}
