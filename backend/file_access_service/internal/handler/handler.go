package handler 

import(
	"log/slog"
	"net/http"
	
	"github.com/go-chi/chi/v5"

	"file-access-service/internal/repository"
)

type Handler struct{
	lgr *slog.Logger
	repo *repository.Repository
}

func NewHandler(lgr *slog.Logger, repo *repository.Repository) *Handler{
	return &Handler{
		lgr: lgr,
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
	w.Write([]byte("addFile"))
}

func (h *Handler) addAccess(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("addAccess"))
}

func (h *Handler) checkAccess(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("checkAccess"))
}