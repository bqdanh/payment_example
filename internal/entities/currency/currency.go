package currency

import (
	"fmt"
	"sync"
)

type Unit string

type Amount struct {
	Currency        Unit  `json:"currency"`
	Amount          int64 `json:"amount"`
	FractionalUnits int64 `json:"fractional_units"`
}

type AmountParser func(amount float64) (Amount, error)

var (
	currencyParserMapping = map[Unit]AmountParser{}
	rwLock                = sync.RWMutex{}
)

func RegisterCurrencyParser(currency Unit, parser AmountParser) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if _, ok := currencyParserMapping[currency]; ok {
		panic(fmt.Errorf("currency(%s) already registered", currency))
	}
	currencyParserMapping[currency] = parser
}

func GetCurrencyParser(currency Unit) (AmountParser, error) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if parser, ok := currencyParserMapping[currency]; ok {
		return parser, nil
	}
	return nil, fmt.Errorf("currency(%s) not found", currency)
}

func FromFloat64(amount float64, currency Unit) (Amount, error) {
	parser, err := GetCurrencyParser(currency)
	if err != nil {
		return Amount{}, fmt.Errorf("get currency(%s) parser: %w", currency, err)
	}
	v, err := parser(amount)
	if err != nil {
		return Amount{}, fmt.Errorf("parse currency(%s) amount: %w", currency, err)
	}
	return v, nil
}
