package controller

import (
	"context"
	"ms-gateway/helper"
	"ms-gateway/pb"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentGRPC pb.PaymentServiceClient
	bookingGRPC pb.BookingServiceClient
}

func NewPaymentController(paymentGRPC pb.PaymentServiceClient, bookingGRPC pb.BookingServiceClient) *PaymentController {
	return &PaymentController{
		paymentGRPC: paymentGRPC,
		bookingGRPC: bookingGRPC,
	}
}

func (p *PaymentController) CreatePayment(c echo.Context) error {
	bookingID := c.Param("booking_id")
	if bookingID == "" {
		return c.JSON(400, helper.Response{
			Message: "booking_id is required",
		})
	}

	in := &pb.GetBookingRequest{
		BookingId: bookingID,
	}

	response, err := p.bookingGRPC.GetBooking(context.TODO(), in)
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
			Message: "you are not authorized to this booking",
		})
	}

	req := new(pb.CreatePaymentRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request",
		})
	}

	req.BookingId = bookingID

	if req.Amount == 0 || req.Method == "" {
		return c.JSON(400, helper.Response{
			Message: "amount and method are required",
		})
	}

	resp, err := p.paymentGRPC.CreatePayment(context.Background(), req)
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to create payment",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, helper.Response{
		Message: "payment created successfully",
		Detail:  resp,
	})
}

func (p *PaymentController) GetPayment(c echo.Context) error {
	paymentID := c.Param("payment_id")
	if paymentID == "" {
		return c.JSON(400, helper.Response{
			Message: "payments_id is required",
		})
	}

	in := &pb.GetPaymentByIdRequest{
		PaymentId: paymentID,
	}

	response, err := p.paymentGRPC.GetPaymentById(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get payment",
			Detail:  err.Error(),
		})
	}

	in2 := &pb.GetBookingRequest{
		BookingId: response.Payment.BookingId,
	}

	checkCredential, err := p.bookingGRPC.GetBooking(context.TODO(), in2)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get booking",
			Detail:  err.Error(),
		})
	}

	userStrCon := c.Get("id").(string)
	userID, _ := strconv.Atoi(userStrCon)

	if response.Payment.BookingId != checkCredential.Booking.BookingId && int32(userID) != checkCredential.Booking.UserId {
		return c.JSON(403, helper.Response{
			Message: "you are not authorized to access this payment",
		})
	}

	return c.JSON(200, response)
}
