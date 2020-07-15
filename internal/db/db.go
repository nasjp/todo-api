package db

import (
	"context"

	"github.com/nasjp/todo-api/internal/todo"
)

type DB interface {
	PutTODO(ctx context.Context, t *todo.TODO) error
	GetAllTODOs(ctx context.Context) (todo.TODOs, error)
}
