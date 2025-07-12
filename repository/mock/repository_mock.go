package mock

import (
	"context"

	"github.com/etakahashi78/mock-example/model"
)

type RepositoryMock struct {
	FindAllMockFunc func(context.Context) ([]model.User, error)
}

func (r *RepositoryMock) FindAll(ctx context.Context) ([]model.User, error) {
	if r.FindAllMockFunc != nil {
		return r.FindAllMockFunc(ctx)
	}
	return nil, nil
}
