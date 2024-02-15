package controller

import (
	"context"
	"ms-booking/model"
	"ms-booking/pb"
	"ms-booking/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingController struct {
	pb.UnimplementedBookingServiceServer
	bookingRepository *repository.BookingRepository
}

func NewBookingController(bookingRepository *repository.BookingRepository) *BookingController {
	return &BookingController{
		bookingRepository: bookingRepository,
	}
}

func (b *BookingController) CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	booking := &model.Booking{
		UserID:       in.UserId,
		RoomID:       in.RoomId,
		CheckinDate:  in.CheckinDate,
		CheckoutDate: in.CheckoutDate,
		Status:       "pending",
	}

	bookingID, err := b.bookingRepository.CreateBooking(booking)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBookingResponse{
		BookingId: bookingID,
	}, nil
}

func (b *BookingController) GetBooking(ctx context.Context, in *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	booking, err := b.bookingRepository.GetBookingByID(in.BookingId)
	if err != nil {
		return nil, err
	}

	bookingProto := &pb.Booking{
		BookingId:    booking.ID.Hex(),
		UserId:       booking.UserID,
		RoomId:       booking.RoomID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
		CreatedAt:    booking.CreatedAt,
		UpdatedAt:    booking.UpdatedAt,
	}

	return &pb.GetBookingResponse{
		Booking: bookingProto,
	}, nil
}

func (b *BookingController) UpdateBooking(ctx context.Context, in *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	bookingID, _ := primitive.ObjectIDFromHex(in.Booking.BookingId)

	booking := &model.Booking{
		ID:           bookingID,
		UserID:       in.Booking.UserId,
		RoomID:       in.Booking.RoomId,
		CheckinDate:  in.Booking.CheckinDate,
		CheckoutDate: in.Booking.CheckoutDate,
		Status:       in.Booking.Status,
		CreatedAt:    in.Booking.CreatedAt,
		UpdatedAt:    time.Now().Format(time.RFC3339),
	}

	err := b.bookingRepository.UpdateBooking(in.Booking.BookingId, booking)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBookingResponse{
		BookingId: booking.ID.Hex(),
	}, nil
}

func (b *BookingController) DeleteBooking(ctx context.Context, in *pb.DeleteBookingRequest) (*pb.DeleteBookingResponse, error) {
	err := b.bookingRepository.DeleteBooking(in.BookingId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteBookingResponse{
		BookingId: in.BookingId,
	}, nil
}

func (b *BookingController) ListBookings(ctx context.Context, in *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	bookings, err := b.bookingRepository.ListBookings(in.UserId)
	if err != nil {
		return nil, err
	}

	var bookingProtos []*pb.Booking
	for _, booking := range bookings {
		bookingProto := &pb.Booking{
			BookingId:    booking.ID.Hex(),
			UserId:       booking.UserID,
			RoomId:       booking.RoomID,
			CheckinDate:  booking.CheckinDate,
			CheckoutDate: booking.CheckoutDate,
			Status:       booking.Status,
			CreatedAt:    booking.CreatedAt,
			UpdatedAt:    booking.UpdatedAt,
		}
		bookingProtos = append(bookingProtos, bookingProto)
	}

	return &pb.ListBookingsResponse{
		Bookings: bookingProtos,
	}, nil
}
