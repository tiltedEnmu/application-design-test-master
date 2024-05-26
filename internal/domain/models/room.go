package models

import "time"

type Room struct {
	HotelID    string      `json:"hotel_id"`
	RoomID     string      `json:"room_id"`
	BookedDays []BookedDay `json:"booked_days"`
}

type BookedDay struct {
	Day        time.Time
	GuestEmail string
}
