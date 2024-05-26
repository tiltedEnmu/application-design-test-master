package main

import (
	"log/slog"
	"os"

	"applicationDesignTest/internal/app"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	application := app.New(log, "localhost", 3000)

	err := application.HttpApp.Run()
	if err != nil {
		panic(err)
	}
}
