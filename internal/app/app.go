package app

import (
	"log/slog"
	"strconv"

	httpapp "applicationDesignTest/internal/app/http"
	"applicationDesignTest/internal/service/booking"
	"applicationDesignTest/internal/storage/inmemory"
)

type App struct {
	HttpApp *httpapp.App
}

func New(
	log *slog.Logger,
	httpHost string, httpPort uint16,
) *App {
	storage := inmemory.New()

	bookingService := booking.New(log, storage, storage)

	httpApp := httpapp.New(log, httpHost, strconv.Itoa(int(httpPort)), bookingService)

	return &App{
		HttpApp: httpApp,
	}
}
