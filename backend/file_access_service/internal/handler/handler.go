package handler 

import(
	"log/slog"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"file-access-service/internal/repository"
	"file-access-service/internal/validator"
)

type Handler struct{
	lgr *slog.Logger
	validator   *validator.CustomValidator
	repo *repository.Repository
}

func NewHandler(lgr *slog.Logger, validator   *validator.CustomValidator, repo *repository.Repository) *Handler{
	return &Handler{
		lgr: lgr,
		validator: validator,
		repo: repo,
	}
}

func (h *Handler) InitRoutes(r *chi.Mux) *chi.Mux {
	r.Route("/access", func(r chi.Router) {
		r.Post("/file", h.addFile)
		r.Post("/file/{fileID}/user", h.addAccess)
		r.Get("/file/{fileID}/user/{userID}", h.checkAccess)
		r.Get("/file/{fileID}", h.getAccessList)
	})
	return r
}

func (h *Handler) addFile(w http.ResponseWriter, r *http.Request){
	const pth = "handler.addFile"
	var req AddFileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithValidationError(w, pth, err, h.lgr)
		return
	}

	if err := h.validator.Validate(&req); err != nil {
		respondWithValidationError(w, pth, validator.GetValidationErrors(err), h.lgr)
		return
	}

	if err := h.repo.CreateFile(r.Context(), req.FileID, req.UserID); err != nil{
		//TODO if already exists
		respondWithError(w, http.StatusInternalServerError, pth, err, ErrorResp{Error: "internal server error"}, h.lgr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) addAccess(w http.ResponseWriter, r *http.Request){
	const pth = "handler.addAccess"
	var req AddAccess
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithValidationError(w, pth, err, h.lgr)
		return
	}

	if err := h.validator.Validate(&req); err != nil {
		respondWithValidationError(w, pth, validator.GetValidationErrors(err), h.lgr)
		return
	}

	if err := h.repo.AddAccess(r.Context(), req.FileID, req.UserID); err != nil{
		respondWithError(w, http.StatusInternalServerError, pth, err, ErrorResp{Error: "internal server error"}, h.lgr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) checkAccess(w http.ResponseWriter, r *http.Request){
	const pth = "handler.team.checkAccess"
	fileID := chi.URLParam(r, "fileID")
	userID := chi.URLParam(r, "userID")
	if fileID == "" || userID == ""{
		respondWithValidationError(w, pth, fmt.Errorf("miss fileID or userID in query"), h.lgr)
		return
	}

	if err := h.repo.CheckAccess(r.Context(), fileID, userID); err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			respondWithError(w, http.StatusForbidden, pth, err, ErrorResp{Error: "user does not have access"}, h.lgr)
			return
		}
		respondWithError(w, http.StatusInternalServerError, pth, err, ErrorResp{Error: "internal server error"}, h.lgr)
		return
	}
	
	resp := AddAccess{FileID: fileID, UserID: userID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) getAccessList(w http.ResponseWriter, r *http.Request){
	const pth = "handler.team.getAccessList"
	fileID := chi.URLParam(r, "fileID")
	if fileID == ""{
		respondWithValidationError(w, pth, fmt.Errorf("miss fileID in query"), h.lgr)
		return
	}

	users, err := h.repo.GetAccessList(r.Context(), fileID)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, pth, err, ErrorResp{Error: "internal server error"}, h.lgr)
		return
	}
	
	resp := AccessList{FileID: fileID, Users: users}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func respondWithError(w http.ResponseWriter, statusCode int, pth string, err error, resp ErrorResp, lgr *slog.Logger) {
	lgr.Error(err.Error(), "handler", pth)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(resp)
}

func respondWithValidationError(w http.ResponseWriter, pth string, err error, lgr *slog.Logger) {
	resp := ErrorResp{
		Error: "validation error",
	}
	lgr.Error(err.Error(), "handler", pth)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	_ = json.NewEncoder(w).Encode(resp)
}
