package db

import (
	"context"
	"sync"

	"github.com/nasjp/todo-api/internal/todo"
)

var _ DB = (*memoryDB)(nil)

func NewMemoryDB(db map[string]*todo.TODO) DB {
	if db == nil {
		db = map[string]*todo.TODO{}
	}

	return &memoryDB{db: db}
}

type memoryDB struct {
	db   map[string]*todo.TODO
	lock sync.RWMutex
}

func (m *memoryDB) PutTODO(ctx context.Context, t *todo.TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()

	return nil
}

func (m *memoryDB) GetAllTODOs(ctx context.Context) (todo.TODOs, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make(todo.TODOs, 0, len(m.db))
	for _, t := range m.db {
		result = append(result, t)
	}

	return result, nil
}
