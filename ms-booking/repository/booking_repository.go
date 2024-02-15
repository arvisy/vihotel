package repository

import (
	"context"
	"errors"
	"ms-booking/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingRepository struct {
	DB *mongo.Database
}

func NewBookingRepository(DB *mongo.Database) *BookingRepository {
	return &BookingRepository{DB: DB}
}

func (b *BookingRepository) CreateBooking(booking *model.Booking) (string, error) {
	booking.CreatedAt = time.Now().Format(time.RFC3339)
	booking.UpdatedAt = time.Now().Format(time.RFC3339)
	result, err := b.DB.Collection("bookings").InsertOne(context.Background(), booking)
	if err != nil {
		return "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to convert InsertedID to primitive.ObjectID")
	}
	return insertedID.Hex(), nil
}

func (b *BookingRepository) GetBookingByID(bookingID string) (*model.Booking, error) {
	id, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return nil, err
	}
	var booking model.Booking
	err = b.DB.Collection("bookings").FindOne(context.Background(), bson.M{"_id": id}).Decode(&booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (b *BookingRepository) UpdateBooking(bookingID string, booking *model.Booking) error {
	id, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return err
	}
	booking.UpdatedAt = time.Now().Format(time.RFC3339)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": booking}
	_, err = b.DB.Collection("bookings").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookingRepository) DeleteBooking(bookingID string) error {
	id, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	_, err = b.DB.Collection("bookings").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookingRepository) ListBookings(userID string) ([]*model.Booking, error) {
	var bookings []*model.Booking
	cursor, err := b.DB.Collection("bookings").Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var booking model.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}