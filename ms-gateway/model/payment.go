package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BookingID  primitive.ObjectID `bson:"booking_id,omitempty"`
	InvoiceID  string             `bson:"invoice_id,omitempty" json:"invoice_id"`
	InvoiceURL string             `bson:"invoice_url,omitempty" json:"invoice_url"`
	Amount     int32              `bson:"amount,omitempty" json:"amount"`
	Method     string             `bson:"method,omitempty" json:"method"`
	Status     string             `bson:"status,omitempty" json:"status"`
	CreatedAt  string             `bson:"created_at,omitempty"`
	UpdatedAt  string             `bson:"updated_at,omitempty"`
}

type CreatePaymentRequest struct {
	Amount int32  `bson:"amount,omitempty" json:"amount"`
	Method string `bson:"method,omitempty" json:"method"`
}
