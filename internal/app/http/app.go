package http

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"applicationDesignTest/internal/transport/http/booking"
)

type App struct {
	log *slog.Logger
	srv *http.Server
}

func New(log *slog.Logger, host string, port string, bookingService booking.Booking) *App {
	mux := http.NewServeMux()
	booking.Register(mux, bookingService)

	return &App{
		log: log,
		srv: &http.Server{
			Addr:           net.JoinHostPort(host, port),
			Handler:        mux,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    time.Second * 5,
			WriteTimeout:   time.Second * 5,
		},
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("addr", a.srv.Addr),
	)

	log.Info("http server is running")

	err := a.srv.ListenAndServe()

	return fmt.Errorf("%s: %w", op, err)
}

// Stop stops http server
func (a *App) Stop(ctx context.Context) error {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping http server")

	err := a.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
