package stripe

import "fmt"

// Внешняя библиотека код которой мы не можем менять

type StripeClient struct {
	APIKey string
}

type PaymentParams struct {
	Amount   int    // сумма в центах
	Currency string // валюта в нижнем регистре ("usd")
	Metadata map[string]string
}

type PaymentResponse struct {
	ID     string
	Status string
	Amount int
}

// метод принимает параметры в другом формате
func (s *StripeClient) CreatePayment(params *PaymentParams) (*PaymentResponse, error) {
	// всякая сложная логика

	if params.Amount > 1000*100 && params.Currency == "usd" { // больше 1000 долларов не принимаем
		return nil, fmt.Errorf("too big price")
	}
	return &PaymentResponse{
		ID:     "42",
		Status: "Success",
		Amount: params.Amount,
	}, nil
}
