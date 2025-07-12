package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/etakahashi78/mock-example/databases"
	"github.com/etakahashi78/mock-example/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func main() {
	db := databases.NewSqlx()
	repo := repository.NewMySQLRepository(db)

	ctx := context.Background()
	if err := run(ctx, repo); err != nil {
		slog.Error("run func failed.", "err", err)
	}

}

func run(ctx context.Context, repo repository.Repository) error {
	s := NewService(repo)

	users, err := s.repo.FindAll(ctx)
	if err != nil {
		slog.Error("FindAll failed.", "err", err)
		return fmt.Errorf("err=%w", err)
	}

	slog.Info("selected rows", "len(users)", len(users))

	return nil
}
