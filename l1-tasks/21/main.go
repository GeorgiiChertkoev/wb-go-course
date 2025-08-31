package main

import (
	"fmt"
	"payments/stripe"
	"strings"
)

// PaymentProcessor - интерфейс для нашего приложения
type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (string, error)
}

type StripeAdapter struct {
	client *stripe.StripeClient
}

func (adapter *StripeAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	response, err := adapter.client.CreatePayment(&stripe.PaymentParams{
		Amount:   int(amount * 100), // переводим в центы
		Currency: strings.ToLower(currency),
		// metadata у нас пустая
	})
	if err != nil {
		return "", err
	}
	if response == nil {
		return "", fmt.Errorf("stripe response and err is nil")
	}
	return response.ID, err
}

func main() {
	var processor PaymentProcessor

	// Создаем адаптер для stripe
	stripeClient := &stripe.StripeClient{APIKey: "sk_test_123"}
	processor = &StripeAdapter{client: stripeClient}

	// Теперь можно использовать с нашим интерфейсом
	paymentID, err := processor.ProcessPayment(99.99, "USD")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Платеж обработан: %s\n", paymentID)
}

/*
Паттерн адаптер позволяет использовать интерфейсы которые не совместимы с нашим
без переписывания всего кода

Плюсы:
- Не надо переписывать весь код при смене библиотеки или
  при работе с несовместимыми интерфейсами
- Адаптер позволяет инкапсулировать логику работы
- Можно использовать для тестирования

Минусы:
- Усложнение кода потому что должны появляться новые классы и интерфейсы
- Скрывает несовместимости в архитектуре и может быть костылем
- Накладные расходы на производительность (незначительные но всё же)
*/
