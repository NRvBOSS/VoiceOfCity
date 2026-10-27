package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode/utf8"

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

func (h *IssueHandler) Create(w http.ResponseWriter, r *http.Request) {
	type createIssueRequest struct {
		Category    string  `json:"category"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	defer r.Body.Close()

	var req createIssueRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Yanlış JSON məlumatı", http.StatusBadRequest)
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	titleLen := utf8.RuneCountInString(req.Title)
	if titleLen < 5 || titleLen > 100 {
		http.Error(w, "Başlıq 5–100 simvol olmalıdır", http.StatusBadRequest)
		return
	}

	descLen := utf8.RuneCountInString(req.Description)
	if descLen < 10 || descLen > 500 {
		http.Error(w, "Təsvir 10–500 simvol olmalıdır", http.StatusBadRequest)
		return
	}

	switch req.Category {
	case "road", "waste", "streetlight", "sidewalk", "other":
	default:
		http.Error(w, "Yanlış kateqoriya", http.StatusBadRequest)
		return
	}

	if req.Latitude < -90 || req.Latitude > 90 ||
		req.Longitude < -180 || req.Longitude > 180 {
		http.Error(w, "Koordinatlar düzgün deyil", http.StatusBadRequest)
		return
	}

	issue, err := h.queries.CreateIssue(r.Context(), db.CreateIssueParams{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
	})
	if err != nil {
		http.Error(w, "Problem bazada saxlanılmadı", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(issue); err != nil {
		return
	}
}
