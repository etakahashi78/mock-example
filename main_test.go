package main

import (
	"context"
	"errors"
	"testing"

	"github.com/etakahashi78/mock-example/model"
	"github.com/etakahashi78/mock-example/repository"
	"github.com/etakahashi78/mock-example/repository/mock"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		repo    repository.Repository
		wantErr bool
	}{
		{
			name: "Happy path.(正常系)",
			ctx:  context.Background(),
			repo: &mock.RepositoryMock{
				FindAllMockFunc: func(ctx context.Context) ([]model.User, error) {
					return []model.User{
						{ID: 1, Name: "testUser1", Email: "test1@example.com"},
						{ID: 2, Name: "testUser2", Email: "test2@example.com"},
					}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "Negative path.(異常系)",
			ctx:  context.Background(),
			repo: &mock.RepositoryMock{
				FindAllMockFunc: func(ctx context.Context) ([]model.User, error) {
					return nil, errors.New("unexpected mock error")
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.ctx, tt.repo); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
