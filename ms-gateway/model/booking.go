package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       int32              `bson:"user_id,omitempty"`
	RoomID       int32              `bson:"room_id,omitempty" json:"room_Id"`
	CheckinDate  string             `bson:"checkin_date,omitempty" json:"checkin_date"`
	CheckoutDate string             `bson:"checkout_date,omitempty" json:"checkout_date"`
	Status       string             `bson:"status,omitempty" json:"status"`
	CreatedAt    string             `bson:"created_at,omitempty"`
	UpdatedAt    string             `bson:"updated_at,omitempty"`
}

type RoomBooking struct {
	ID       int32 `bson:"_id,omitempty"`
	HotelID  int32 `bson:"hotel_id,omitempty"`
	Capacity int   `bson:"capacity,omitempty"`
	Price    int   `bson:"price,omitempty"`
}

type UserBooking struct {
	ID    int32  `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
}

type CreateBookingRequest struct {
	RoomID       int32       `json:"room_id"`
	CheckinDate  string      `json:"checkin_date"`
	CheckoutDate string      `json:"checkout_date"`
	Room         RoomRequest `json:"room"`
}

type RoomRequest struct {
	HotelID  string `json:"hotel_id"`
	Capacity string `json:"capacity"`
}
