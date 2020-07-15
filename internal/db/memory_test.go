package db_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp/todo-api/internal/db"
	"github.com/nasjp/todo-api/internal/todo"
)

func TestMemoryDBPutTODO(t *testing.T) {
	t.Parallel()

	type arg struct {
		ctx  context.Context
		todo *todo.TODO
	}

	todo1 := &todo.TODO{
		ID:    "75035d56-2a23-4fbe-875e-97652efbb3d1",
		Title: "Buy mac book",
	}

	tests := []struct {
		name string
		arg  arg
		want map[string]*todo.TODO
	}{
		{"put", arg{ctx: context.Background(), todo: todo1}, map[string]*todo.TODO{todo1.ID: todo1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := map[string]*todo.TODO{}
			if err := db.NewMemoryDB(d).PutTODO(tt.arg.ctx, tt.arg.todo); err != nil {
				t.Errorf("Unexpected error in putting todo: %v", err)
				return
			}

			if diff := cmp.Diff(tt.want, d); diff != "" {
				t.Errorf("putting todo results mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestMemoryDBGetAllTODOs(t *testing.T) {
	t.Parallel()

	todo1 := &todo.TODO{
		ID:    "75035d56-2a23-4fbe-875e-97652efbb3d1",
		Title: "Buy mac book",
	}

	todo2 := &todo.TODO{
		ID:    "bcf60a62-fb9b-4002-9f84-f81b9fab92c2",
		Title: "Buy house",
	}

	tests := []struct {
		name string
		arg  context.Context
		want todo.TODOs
	}{
		{"put", context.Background(), todo.TODOs{todo1, todo2}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := map[string]*todo.TODO{
				todo1.ID: todo1,
				todo2.ID: todo2,
			}
			ts, err := db.NewMemoryDB(d).GetAllTODOs(tt.arg)
			if err != nil {
				t.Errorf("Unexpected error in putting todo: %v", err)
				return
			}

			if diff := cmp.Diff(tt.want, ts); diff != "" {
				t.Errorf("putting todo results mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
