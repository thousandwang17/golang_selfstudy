package service

import (
	"context"

	"gokit/pkg/repository"

	"github.com/go-kit/kit/log"
)

type Service interface {
	Put(ctx context.Context, key, val string) (err error)
}
type service struct {
	logger     log.Logger
	repository repository.Repository
}

func (s *service) Put(ctx context.Context, key, val string) (err error) {
	return s.repository.Put(key, val)
}
func New(logger log.Logger, repository repository.Repository) Service {
	return &service{logger: logger, repository: repository}
}
