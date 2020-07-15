package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/nasjp/todo-api/internal/db"
	"github.com/nasjp/todo-api/internal/todo"
)

var _ http.Handler = (*createHandler)(nil)
var _ http.Handler = (*listHandler)(nil)

type createHandler struct {
	db db.DB
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := &todo.TODO{}

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	t.ID = uuid.New().String()

	if err := h.db.PutTODO(r.Context(), t); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(t); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

type listHandler struct {
	db db.DB
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ts, err := h.db.GetAllTODOs(r.Context())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(ts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
