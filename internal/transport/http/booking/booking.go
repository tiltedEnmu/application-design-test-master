package booking

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"applicationDesignTest/internal/domain/models"
)

var (
	ErrInternal           = errors.New("internal error")
	ErrDayAlreadyReserved = errors.New("day already reserved")
)

type Booking interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetBookedRooms(ctx context.Context) ([]models.Room, error)
}

type api struct {
	booking Booking
}

func Register(mux *http.ServeMux, service Booking) {
	a := &api{booking: service}
	mux.HandleFunc("/order", a.createOrder)
	mux.HandleFunc("/rooms", a.getBookedRooms)
}

func (a *api) createOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var newOrder models.Order
		err := json.NewDecoder(r.Body).Decode(&newOrder)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("provided json object is invalid"))
			return
		}

		err = a.booking.CreateOrder(r.Context(), newOrder)
		if err != nil {
			switch {
			case errors.Is(err, ErrDayAlreadyReserved):
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte("one or more of the days provided are already reserved"))
			case errors.Is(err, ErrInternal):
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal error"))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("unknown error"))
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("order processed"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", "POST")
	}
}

func (a *api) getBookedRooms(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rooms, err := a.booking.GetBookedRooms(r.Context())

		if err != nil {
			switch {
			case errors.Is(err, ErrInternal):
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal error"))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("unknown error"))
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if len(rooms) == 0 {
			_, _ = w.Write([]byte("no reserved rooms"))
			return
		}

		err = json.NewEncoder(w).Encode(rooms)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("internal error"))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", "GET")
	}
}
