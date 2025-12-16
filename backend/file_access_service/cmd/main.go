package main

import(
	"os"
	"os/signal"
	"log/slog"
	"net/http"
	"syscall"
	"time"
	"context"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"

	"file-access-service/internal/database"
	"file-access-service/internal/repository"
	"file-access-service/internal/handler"
	"file-access-service/internal/logger"
	"file-access-service/internal/validator"
)

func main(){
	lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	lgr.Info("initializing database")
	db, err := database.New()
	if err != nil {
		lgr.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	err = database.Migrate(db)
	if err != nil {
		lgr.Error("Migration failed: %v", err)
		os.Exit(1)
	}
	lgr.Info("migration completed successfully")

	
	r := repository.NewRepository(db)
	v := validator.NewValidator()
	h := handler.NewHandler(lgr, v, r)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://localhost:3000"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300,
	}))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(logger.New(lgr))

	h.InitRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lgr.Info("starting server ...")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lgr.Error("listen", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	lgr.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		lgr.Error("server shutdown:", "error", err)
		os.Exit(1)
	}
	<-ctx.Done()
	lgr.Info("timeout of 5 seconds.")
	lgr.Info("server exiting")
}