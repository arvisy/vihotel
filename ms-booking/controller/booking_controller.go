package controller

import (
	"context"
	"ms-booking/model"
	"ms-booking/pb"
	"ms-booking/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	booking := &model.BookingDetails{
		User: model.User{
			ID:    in.User.UserId,
			Name:  in.User.Name,
			Email: in.User.Email,
		},
		Room: model.Room{
			ID:       in.Room.RoomId,
			HotelID:  in.Room.HotelId,
			Capacity: int(in.Room.Capacity),
			Price:    int(in.Room.Price),
		},
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

func (s *BookingController) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	// Validasi booking ID
	bookingID, err := primitive.ObjectIDFromHex(req.BookingId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid booking ID: %v", err)
	}

	// Update booking di repository
	if err := s.bookingRepository.UpdateBooking(ctx, bookingID, req.Booking); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update booking: %v", err)
	}

	// Buat respons
	resp := &pb.UpdateBookingResponse{
		BookingId: req.BookingId,
	}

	return resp, nil
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
