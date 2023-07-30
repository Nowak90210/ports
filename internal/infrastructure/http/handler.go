package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Nowak90210/ports/internal/app"
	"github.com/Nowak90210/ports/internal/infrastructure"
	"github.com/gorilla/mux"
)

type apiError struct {
	Error string `json:"error"`
}

type postFileResponse struct {
	SavedRows int `json:"saved_rows"`
}

type Handler struct {
	service *app.Service
}

func NewHandler(service *app.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/port/{id}", h.handleGetPort)
	router.HandleFunc("/file", h.handlePostFile)

	return router
}

func (h *Handler) handleGetPort(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		h.writeJSON(w, http.StatusBadRequest, apiError{Error: "id missing"})
	}

	port, err := h.service.Get(id)
	if err != nil {
		h.writeJSON(w, http.StatusNotFound, apiError{Error: err.Error()})
		return
	}

	dto := infrastructure.FromDomain(port)
	h.writeJSON(w, http.StatusOK, dto)
}

func (h *Handler) handlePostFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		h.writeJSON(w, http.StatusBadRequest, apiError{Error: fmt.Sprintf("unsuported method '%s'", r.Method)})
		return
	}

	fileName := r.FormValue("file_name")
	if fileName == "" {
		h.writeJSON(w, http.StatusBadRequest, apiError{Error: "fileName missing"})
		return
	}

	// preventing path traversal
	cleanFileName := filepath.Clean(fileName)
	counter, err := h.service.SavePortsFromFile(cleanFileName)
	if err != nil {
		h.writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		return
	}

	h.writeJSON(w, http.StatusCreated, postFileResponse{SavedRows: counter})
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
