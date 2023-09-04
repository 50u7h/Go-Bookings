package dbrepo

import (
	"errors"
	"github.com/50u7h/Go-Bookings/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservations insert a reservations into the database
func (m *testDBRepo) InsertReservations(res models.Reservation) (int, error) {
	// if room id is 2 fail; otherwise, pass
	if res.RoomID == 1 {
		return 0, errors.New("Error")
	}
	return 1, nil
}

// InsertRoomRestriction insert a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	if res.RoomID == 1000 {
		return errors.New("Error")
	}
	return nil
}

// SearchAvailabilityByDatesAndRoomID returns true if availability exists for roomID, and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesAndRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRoms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRoms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID get a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("ERROR")
	}
	return room, nil
}
