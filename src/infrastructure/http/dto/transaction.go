package dto

import "time"

type (
	ChargeTransaction struct {
		ExternalId          string              `json:"-"` //Transaction ID that we generate
		MerchantId          string              `json:"-"`
		ClientTransactionID string              `json:"client_transaction_id" validate:"required"`
		Amount              float64             `json:"amount" validate:"required"`
		Description         string              `json:"description"`
		PaymentMethod       string              `json:"payment_method" validate:"required"`
		VirtualAccount      ChargeTransactionVA `json:"virtual_account"`
	}

	ChargeTransactionVA struct {
		Bank string `json:"bank"`
		Name string `json:"name"`
	}

	ChargeResponse struct {
		Id               string           `json:"id"`
		ExternalId       string           `json:"external_id"`
		UserId           string           `json:"user_id"`
		Status           string           `json:"status"`
		Amount           float64          `json:"amount"`
		Description      string           `json:"description"`
		ExpiryDate       *time.Time       `json:"expiry_date"`
		InvoiceUrl       string           `json:"invoice_url"`
		QRString         string           `json:"qr_string"`
		ReferenceId      string           `json:"reference_id"`
		Currency         string           `json:"currency"`
		QRPaymentWebhook QRPaymentWebhook `json:"qr_payment_webhook"`
	}

	QRPaymentWebhook struct {
		Created    time.Time `json:"created"`
		BusinessID string    `json:"business_id"`
		Event      string    `json:"event"`
		APIVersion string    `json:"api_version"`
		Data       QRData    `json:"data"`
	}

	QRData struct {
		Amount        int64         `json:"amount"`
		BusinessID    string        `json:"business_id"`
		ChannelCode   string        `json:"channel_code"`
		Created       time.Time     `json:"created"`
		Currency      string        `json:"currency"`
		ExpiresAt     time.Time     `json:"expires_at"`
		ID            string        `json:"id"`
		Metadata      any           `json:"metadata"`
		PaymentDetail PaymentDetail `json:"payment_detail"`
		QRId          string        `json:"qr_id"`
		QRString      string        `json:"qr_string"`
		ReferenceID   string        `json:"reference_id"`
		Status        string        `json:"status"`
		Type          string        `json:"type"`
	}

	VAPaymentWebhook struct {
		Id                       string      `json:"id"`
		Amount                   int         `json:"amount"`
		Country                  string      `json:"country"`
		Created                  time.Time   `json:"created"`
		Updated                  time.Time   `json:"updated"`
		Currency                 string      `json:"currency"`
		OwnerId                  string      `json:"owner_id"`
		BankCode                 string      `json:"bank_code"`
		PaymentId                string      `json:"payment_id"`
		ExternalId               string      `json:"external_id"`
		MerchantCode             string      `json:"merchant_code"`
		AccountNumber            string      `json:"account_number"`
		PaymentDetail            interface{} `json:"payment_detail"`
		TransactionTimestamp     time.Time   `json:"transaction_timestamp"`
		CallbackVirtualAccountId string      `json:"callback_virtual_account_id"`
	}

	PaymentDetail struct {
		AccountDetails *any    `json:"account_details"`
		CustomerPAN    *string `json:"customer_pan"`
		MerchantPAN    *string `json:"merchant_pan"`
		Name           *string `json:"name"`
		ReceiptID      string  `json:"receipt_id"`
		Source         string  `json:"source"`
	}
)
