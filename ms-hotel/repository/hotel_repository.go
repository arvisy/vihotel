package repository

import (
	"database/sql"
	"ms-hotel/model"
	"time"
)

type HotelRepository struct {
	DB *sql.DB
}

func NewHotelRepository(DB *sql.DB) *HotelRepository {
	return &HotelRepository{DB: DB}
}

func (h *HotelRepository) CreateHotel(name, location, description string) (int32, error) {
	var id int32
	err := h.DB.QueryRow("INSERT INTO hotels (name, location, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", name, location, description, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (h *HotelRepository) UpdateHotel(id int32, name, location, description string) error {
	_, err := h.DB.Exec("UPDATE hotels SET name=$1, location=$2, description=$3, updated_at=$4 WHERE id=$5", name, location, description, time.Now().Format(time.RFC3339), id)
	return err
}

func (h *HotelRepository) DeleteHotel(id int32) error {
	_, err := h.DB.Exec("DELETE FROM hotels WHERE id=$1", id)
	return err
}

func (h *HotelRepository) GetHotel(id int32) (*model.Hotel, error) {
	var hotel model.Hotel
	err := h.DB.QueryRow("SELECT id, name, location, description, created_at, updated_at FROM hotels WHERE id=$1", id).Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Description, &hotel.CreatedAt, &hotel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (h *HotelRepository) ListHotels() ([]*model.Hotel, error) {
	rows, err := h.DB.Query("SELECT id, name, location, description, created_at, updated_at FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []*model.Hotel
	for rows.Next() {
		var hotel model.Hotel
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Description, &hotel.CreatedAt, &hotel.UpdatedAt)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, &hotel)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (h *HotelRepository) CreateRoom(hotelID int, capacity int, price int) (int, error) {
	var id int
	err := h.DB.QueryRow("INSERT INTO rooms (hotel_id, capacity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", hotelID, capacity, price, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (h *HotelRepository) UpdateRoom(id int, hotelID int, capacity int, price int) error {
	_, err := h.DB.Exec("UPDATE rooms SET hotel_id=$1, capacity=$2, price=$3, updated_at=$4 WHERE id=$5", hotelID, capacity, price, time.Now(), id)
	return err
}

func (h *HotelRepository) DeleteRoom(id int) error {
	_, err := h.DB.Exec("DELETE FROM rooms WHERE id=$1", id)
	return err
}

func (h *HotelRepository) GetRoom(id int) (*model.Room, error) {
	var room model.Room
	err := h.DB.QueryRow("SELECT id, hotel_id, capacity, price, created_at, updated_at FROM rooms WHERE id=$1", id).Scan(&room.Id, &room.HotelID, &room.Capacity, &room.Price, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (h *HotelRepository) ListRooms(hotelID int) ([]*model.Room, error) {
	rows, err := h.DB.Query("SELECT id, hotel_id, capacity, price, created_at, updated_at FROM rooms WHERE hotel_id=$1", hotelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.Id, &room.HotelID, &room.Capacity, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
