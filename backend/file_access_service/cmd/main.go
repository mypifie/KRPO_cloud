package main

import(
	"os"
	"log/slog"
	"file-access-service/internal/database"
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
}