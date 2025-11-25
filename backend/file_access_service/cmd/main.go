package main

import(
	"os"
	"fmt"
	"log/slog"
	"file-access-service/internal/database"
)

func main(){
	lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	lgr.Info("initializing database")
	db, err := database.New()
	if err != nil {
		lgr.Error("failed to initialize database", "error", err)
	}
	fmt.Println(db)
}