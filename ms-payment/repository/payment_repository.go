package repository

import (
	"context"
	"ms-payment/client"
	"ms-payment/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	DB *mongo.Database
}

func NewPaymentRepository(DB *mongo.Database) *PaymentRepository {
	return &PaymentRepository{
		DB: DB,
	}
}

func (p *PaymentRepository) CreatePayment(payment *model.Payment) error {
	invoice, err := client.CreateInvoice(payment.BookingID.Hex(), payment.Amount, &payment.Method)
	if err != nil {
		return err
	}

	payment.InvoiceID = *invoice.Id
	payment.InvoiceURL = invoice.InvoiceUrl

	now := time.Now().UTC().Format(time.RFC3339)
	payment.CreatedAt = now
	payment.UpdatedAt = now

	collection := p.DB.Collection("payments")
	_, err = collection.InsertOne(context.Background(), payment)
	if err != nil {
		return err
	}

	bookingsCollection := p.DB.Collection("bookings")
	_, err = bookingsCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": payment.BookingID},
		bson.M{"$set": bson.M{"status": "success"}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentRepository) GetPaymentByID(paymentID string) (*model.Payment, error) {
	id, err := primitive.ObjectIDFromHex(paymentID)
	if err != nil {
		return nil, err
	}
	var payment model.Payment
	err = p.DB.Collection("payments").FindOne(context.Background(), bson.M{"_id": id}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
