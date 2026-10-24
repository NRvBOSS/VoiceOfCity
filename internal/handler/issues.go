package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NRvBOSS/citysesi/internal/db"
)

type IssueHandler struct {
	queries *db.Queries
}

func NewIssueHandler(queries *db.Queries) *IssueHandler {
	return &IssueHandler{
		queries: queries,
	}
}

func (h *IssueHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := db.ListIssuesParams{
		Limit:  20,
		Offset: 0,
	}

	issues, err := h.queries.ListIssues(ctx, params)
	if err != nil {
		http.Error(w, "Problemləri əldə etmək mümkün olmadı", http.StatusInternalServerError)
		return
	}

	if issues == nil {
		issues = []db.Issue{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(issues); err != nil {
		return
	}
}
