package inmemory

import (
	"context"
	"strings"
	"time"

	"applicationDesignTest/internal/domain/models"
)

func (s *Storage) ReserveRoom(
	ctx context.Context,
	hotelID, roomID string,
	guestEmail string,
	daysForBooking []time.Time,
) error {
	const op = "storage.inmemory.ReserveRoom"

	s.mu.Lock()
	defer s.mu.Unlock()

	days := make([]models.BookedDay, len(daysForBooking))
	for k, v := range daysForBooking {
		days[k] = models.BookedDay{Day: v, GuestEmail: guestEmail}
	}

	s.rooms[hotelID+":"+roomID] = append(s.rooms[hotelID+":"+roomID], days...)

	return nil
}

func (s *Storage) CheckForIntersections(
	ctx context.Context,
	hotelID, roomID string,
	daysForBooking []time.Time,
) ([]time.Time, error) {
	const op = "storage.inmemory.CheckForIntersections"

	s.mu.RLock()
	defer s.mu.RUnlock()

	room, ok := s.rooms[hotelID+":"+roomID]
	if !ok {
		return nil, nil
	}

	var unavailableDays []time.Time

	for _, v := range daysForBooking {
		for _, j := range room {
			if v == j.Day {
				unavailableDays = append(unavailableDays, v)
				break
			}
		}
	}

	return unavailableDays, nil
}

func (s *Storage) AllHotelsRooms(
	ctx context.Context,
) ([]models.Room, error) {
	const op = "storage.inmemory.Room"

	s.mu.RLock()
	defer s.mu.RUnlock()

	var rooms []models.Room
	for k, v := range s.rooms { // k - hotel room ID, v = already booked days for current room
		id := strings.Split(k, ":")
		rooms = append(rooms, models.Room{HotelID: id[0], RoomID: id[1], BookedDays: v})
	}

	return rooms, nil
}
