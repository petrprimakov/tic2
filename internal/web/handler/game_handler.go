package handler

import (
	"encoding/json"

	"errors"
	"net/http"
	"tic2/internal/domain/service"
	apperrors "tic2/internal/errors"
	webmodel "tic2/internal/web/model"

	"github.com/google/uuid"
)

type GameHandler struct {
	svc service.GameService
}

func NewGameHandler(svc service.GameService) *GameHandler {
	return &GameHandler{svc: svc}
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	game, err := h.svc.CreateGame(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, webmodel.GameResponseFromDomain(game, false, 0))
}

func (h *GameHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /game", h.CreateGame)
	mux.HandleFunc("POST /game/{id}", h.PostMove)
	mux.HandleFunc("DELETE /game/{id}", h.DeleteGame)
}

func (h *GameHandler) PostMove(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid game UUID: "+err.Error())
		return
	}

	var req webmodel.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "malformed request body: "+err.Error())
		return
	}

	game := req.ToDomain(id)

	updated, err := h.svc.ComputeNextMove(r.Context(), game)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	finished, winner := h.svc.CheckGameOver(r.Context(), updated.Board)
	writeJSON(w, http.StatusOK, webmodel.GameResponseFromDomain(updated, finished, winner))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, webmodel.ErrorResponse{Error: msg})
}

func (h *GameHandler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid game UUID: "+err.Error())
		return
	}

	if err := h.svc.DeleteGame(r.Context(), id); err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
