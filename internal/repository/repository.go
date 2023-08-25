package repository

import "github.com/50u7h/Go-Bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservations(res models.Reservation) error
}
