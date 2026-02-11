package wallet

import "time"

type TransactionLog struct {
	ID                            string    `json:"id,omitempty"`
	TrxID                         string    `json:"trx_id,omitempty"`
	PaymentGatewayPayload         []byte    `json:"payment_gateway_payload,omitempty"`
	PaymentGatewayResponse        []byte    `json:"payment_gateway_response,omitempty"`
	InfoLevel                     string    `json:"info_level,omitempty"`
	RequestPayload                []byte    `json:"payload,omitempty"`
	PaymentGatewayCallbackPayload []byte    `json:"payment_gateway_callback_payload,omitempty"`
	CreatedAt                     time.Time `json:"created_at,omitempty"`
	UpdatedAt                     time.Time `json:"updated_at,omitempty"`
}

const (
	TransactionLogLevelInfo  = "info"
	TransactionLogLevelWarn  = "warn"
	TransactionLogLevelError = "error"
)
