package controller

import "ms-gateway/pb"

type BookingController struct {
	bookingGRPC pb.BookingServiceClient
}

func NewBookingController(bookingGRPC pb.BookingServiceClient) *BookingController {
	return &BookingController{
		bookingGRPC: bookingGRPC,
	}
}
