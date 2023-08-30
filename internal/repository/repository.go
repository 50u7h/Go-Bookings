package repository

import (
	"github.com/50u7h/Go-Bookings/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservations(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesAndRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRoms(start, end time.Time) ([]models.Room, error)
}
