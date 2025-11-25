package handler 

import(
	"log/slog"
	"net/http"
	"encoding/json"
	
	"github.com/go-chi/chi/v5"

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
	w.Write([]byte("addAccess"))
}

func (h *Handler) checkAccess(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("checkAccess"))
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
