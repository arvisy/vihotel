package controller

import (
	"context"
	"ms-payment/model"
	"ms-payment/pb"
	"ms-payment/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentController struct {
	pb.UnimplementedPaymentServiceServer
	paymentRepository *repository.PaymentRepository
}

func NewPaymentController(paymentRepository *repository.PaymentRepository) *PaymentController {
	return &PaymentController{
		paymentRepository: paymentRepository,
	}
}

func (p *PaymentController) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	bookingID, _ := primitive.ObjectIDFromHex(in.BookingId)

	payment := &model.Payment{
		BookingID: bookingID,
		Method:    in.Method,
		Amount:    in.Amount,
		Status:    "success",
	}

	err := p.paymentRepository.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	response := &pb.CreatePaymentResponse{
		BookingId:  payment.BookingID.Hex(),
		InvoiceId:  payment.InvoiceID,
		InvoiceUrl: payment.InvoiceURL,
		Amount:     payment.Amount,
		Method:     payment.Method,
		Status:     payment.Status,
		CreatedAt:  payment.CreatedAt,
		UpdatedAt:  payment.UpdatedAt,
	}

	return response, nil
}

func (p *PaymentController) GetPaymentById(ctx context.Context, in *pb.GetPaymentByIdRequest) (*pb.GetPaymentByIdResponse, error) {
	payment, err := p.paymentRepository.GetPaymentByID(in.PaymentId)
	if err != nil {
		return nil, err
	}

	paymentProto := &pb.Payment{
		Id:         payment.ID.Hex(),
		BookingId:  payment.ID.Hex(),
		InvoiceId:  payment.InvoiceID,
		InvoiceUrl: payment.InvoiceURL,
		Amount:     payment.Amount,
		Method:     payment.Method,
		Status:     payment.Status,
		CreatedAt:  payment.CreatedAt,
		UpdatedAt:  payment.UpdatedAt,
	}

	return &pb.GetPaymentByIdResponse{
		Payment: paymentProto,
	}, nil
}
