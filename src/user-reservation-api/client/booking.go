package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"user-reservation-api/model"
)

// Db es una variable global que representa la conexión a la base de datos.
// var Db *gorm.DB

// GetBookingById obtiene una reserva por su ID, incluyendo información sobre el usuario asociado.
func GetBookingById(id int) model.Booking {
	var booking model.Booking

	if Db == nil {
		// Si la conexión a la base de datos no está disponible, se retorna una reserva vacía.
		return booking
	}

	Db.Where("id = ?", id).Preload("User").First(&booking)
	log.Debug("Booking: ", booking)

	return booking
}

// GetBookings obtiene todas las reservas con su información asociada.
func GetBookings() model.Bookings {
	var bookings model.Bookings
	Db.Find(&bookings)

	log.Debug("Bookings: ", bookings)

	return bookings
}

// InsertBooking inserta una nueva reserva en la base de datos.
func InsertBooking(booking model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		// TODO: gestionar los errores de manera más robusta.
		log.Error("")
	}

	log.Debug("Booking Created: ", booking.Id)
	return booking
}

// GetAvailabilityByIdAndDate verifica la disponibilidad de un hotel en una fecha específica.
func GetAvailabilityByIdAndDate(id_hotel int, startDate int) bool {
	var booking model.Booking

	// Imprime un mensaje antes y después de la consulta a la base de datos (puede ser útil para depuración).
	fmt.Println("Antes de la consulta a la base de datos")
	result := Db.Where("hotel_id = ? AND start_date <= ? AND end_date > ?", id_hotel, startDate, startDate).First(&booking) // hay reserva aquí!
	fmt.Println("Después de la consulta a la base de datos")

	if result.Error != nil {
		return false // No hay reserva, por lo tanto, está disponible.
	}

	return true // Hay una reserva, por lo tanto, no está disponible.
}

// GetBookingByUserId obtiene una reserva por el ID del usuario, incluyendo información sobre el usuario asociado.
func GetBookingByUserId(id int) model.Booking {
	var booking model.Booking

	Db.Where("user_id = ?", id).Preload("User").First(&booking)
	log.Debug("Booking: ", booking)

	return booking
}
