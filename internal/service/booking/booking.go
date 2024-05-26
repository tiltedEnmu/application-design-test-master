package booking

import (
	"context"
	"log/slog"
	"time"

	"applicationDesignTest/internal/domain/models"
	"applicationDesignTest/internal/domain/tools"
	"applicationDesignTest/internal/transport/http/booking"
)

type Booking struct {
	log          *slog.Logger
	roomSaver    RoomSaver
	roomProvider RoomProvider
}

func New(log *slog.Logger, saver RoomSaver, provider RoomProvider) *Booking {
	return &Booking{
		log:          log,
		roomSaver:    saver,
		roomProvider: provider,
	}
}

type RoomSaver interface {
	ReserveRoom(
		ctx context.Context,
		hotelID, roomID string,
		guestEmail string,
		daysForBooking []time.Time,
	) error
}

type RoomProvider interface {
	CheckForIntersections(
		ctx context.Context,
		hotelID, roomID string,
		daysForBooking []time.Time,
	) ([]time.Time, error)
	AllHotelsRooms(
		ctx context.Context,
	) ([]models.Room, error)
}

func (s *Booking) CreateOrder(ctx context.Context, order models.Order) error {
	const op = "service.booking.CreateOrder"

	log := s.log.With(
		slog.String("op", op),
		slog.String("room", order.HotelID+":"+order.RoomID),
	)

	log.Info("performing user's order")

	daysToBook := tools.DaysBetween(order.From, order.To)

	unavailableDates, err := s.roomProvider.CheckForIntersections(ctx, order.HotelID, order.RoomID, daysToBook)
	if err != nil {
		log.Error(err.Error())
		return booking.ErrInternal
	}

	if len(unavailableDates) != 0 {
		log.Error("provided days already reserved", slog.Any("dates", unavailableDates))
		return booking.ErrDayAlreadyReserved
	}

	err = s.roomSaver.ReserveRoom(ctx, order.HotelID, order.RoomID, order.UserEmail, daysToBook)
	if err != nil {
		log.Error(err.Error())
		return booking.ErrInternal
	}

	return nil
}

func (s *Booking) GetBookedRooms(ctx context.Context) ([]models.Room, error) {
	const op = "service.booking.GetRoom"

	log := s.log.With(slog.String("op", op))

	rooms, err := s.roomProvider.AllHotelsRooms(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil, booking.ErrInternal
	}

	return rooms, nil
}
