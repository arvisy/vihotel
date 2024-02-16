package controller

import (
	"context"
	"ms-gateway/helper"
	"ms-gateway/model"
	"ms-gateway/pb"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	userGRPC    pb.UserServiceClient
	hotelGRPC   pb.HotelServiceClient
	bookingGRPC pb.BookingServiceClient
}

func NewBookingController(userGRPC pb.UserServiceClient, hotelGRPC pb.HotelServiceClient, bookingGRPC pb.BookingServiceClient) *BookingController {
	return &BookingController{
		userGRPC:    userGRPC,
		hotelGRPC:   hotelGRPC,
		bookingGRPC: bookingGRPC,
	}
}

// @Summary      Private
// @Description  Create booking
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param		 data body model.CreateBookingRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/booking [Post]
func (b *BookingController) CreateBooking(c echo.Context) error {
	userStrCon := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrCon)

	var payload model.Booking

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.CheckinDate == "" || payload.CheckoutDate == "" {
		return c.JSON(400, helper.Response{
			Message: "room_id, checkin_date, and checkout_date are required",
		})
	}

	userInfo, err := b.userGRPC.GetCustomer(context.Background(), &pb.GetCustomerRequest{Id: int32(userID)})
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get user information",
			Detail:  err.Error(),
		})
	}

	roomInfo, err := b.hotelGRPC.GetRoom(context.Background(), &pb.GetRoomRequest{Id: payload.RoomID})
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get room information",
			Detail:  err.Error(),
		})
	}

	in := &pb.CreateBookingRequest{
		User: &pb.UserInfo{
			UserId: int32(userID),
			Name:   userInfo.Name,
			Email:  userInfo.Email,
		},
		Room: &pb.RoomInfo{
			RoomId:   payload.RoomID,
			HotelId:  roomInfo.Room.HotelId,
			Capacity: roomInfo.Room.Capacity,
			Price:    roomInfo.Room.Price,
		},
		CheckinDate:  payload.CheckinDate,
		CheckoutDate: payload.CheckoutDate,
	}

	response, err := b.bookingGRPC.CreateBooking(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to create booking",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, response)
}

// @Summary      Private
// @Description  Get booking
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Success      201  {object}  model.Booking
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/booking/:booking_id [Get]
func (b *BookingController) GetBooking(c echo.Context) error {
	bookingID := c.Param("booking_id")
	if bookingID == "" {
		return c.JSON(400, helper.Response{
			Message: "booking_id is required",
		})
	}

	in := &pb.GetBookingRequest{
		BookingId: bookingID,
	}

	response, err := b.bookingGRPC.GetBooking(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get booking",
			Detail:  err.Error(),
		})
	}

	userStrCon := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrCon)

	if response.Booking.UserId != int32(userID) {
		return c.JSON(403, helper.Response{
			Message: "you are not authorized to access this booking",
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Update booking
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param		 data body model.CreateBookingRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/booking/:booking_id [Put]
func (b *BookingController) UpdateBooking(c echo.Context) error {
	userStrCon := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrCon)

	bookingID := c.Param("booking_id")
	if bookingID == "" {
		return c.JSON(400, helper.Response{
			Message: "booking ID is required",
		})
	}

	var payload model.Booking
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.RoomID == 0 || payload.CheckinDate == "" || payload.CheckoutDate == "" {
		return c.JSON(400, helper.Response{
			Message: "room ID, check-in date, and check-out date are required",
		})
	}

	checkCredentialBookingID := &pb.GetBookingRequest{
		BookingId: bookingID,
	}

	CredentialBookingID, err := b.bookingGRPC.GetBooking(context.TODO(), checkCredentialBookingID)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get booking",
			Detail:  err.Error(),
		})
	}

	if CredentialBookingID.Booking.UserId != int32(userID) {
		return c.JSON(403, helper.Response{
			Message: "you are not authorized to update this booking",
		})
	}

	in := &pb.UpdateBookingRequest{
		BookingId: bookingID,
		Booking: &pb.Booking{
			UserId:       int32(userID),
			RoomId:       payload.RoomID,
			CheckinDate:  payload.CheckinDate,
			CheckoutDate: payload.CheckoutDate,
			Status:       payload.Status,
		},
	}

	response, err := b.bookingGRPC.UpdateBooking(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to update booking",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Update booking
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/booking/:booking_id [Delete]
func (b *BookingController) DeleteBooking(c echo.Context) error {
	bookingID := c.Param("booking_id")
	if bookingID == "" {
		return c.JSON(400, helper.Response{
			Message: "bookingID is required",
		})
	}

	userStrCon := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrCon)

	checkCredentialBookingID := &pb.GetBookingRequest{
		BookingId: bookingID,
	}

	CredentialBookingID, err := b.bookingGRPC.GetBooking(context.TODO(), checkCredentialBookingID)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get booking",
			Detail:  err.Error(),
		})
	}

	if CredentialBookingID.Booking.UserId != int32(userID) {
		return c.JSON(403, helper.Response{
			Message: "you are not authorized to update this booking",
		})
	}

	in := &pb.DeleteBookingRequest{
		BookingId: bookingID,
	}

	response, err := b.bookingGRPC.DeleteBooking(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to delete booking",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Get all booking
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Success      201  {object}  model.Booking
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/booking [Get]
func (b *BookingController) ListBookings(c echo.Context) error {
	userStrConv := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrConv)

	if userID == 0 {
		return c.JSON(400, helper.Response{
			Message: "userID is required",
		})
	}

	in := &pb.ListBookingsRequest{
		UserId: int32(userID),
	}

	response, err := b.bookingGRPC.ListBookings(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get bookings",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}
