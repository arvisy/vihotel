package client

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/invoice"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func CreateInvoice(bookingID string, amount int32, method *string) (*invoice.Invoice, error) {
	LoadEnv()

	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(bookingID, float64(amount))
	createInvoiceRequest.SetCurrency("IDR")
	if method != nil {
		createInvoiceRequest.SetPaymentMethods([]string{*method})
	}

	xenditClient := xendit.NewClient(os.Getenv("XENDIT_API_KEY"))

	resp, _, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
