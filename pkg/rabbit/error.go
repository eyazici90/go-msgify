package rabbit

import "errors"

var (
	ErrNoChannelFound error = errors.New("no channel found. Please connect to exchange first")

	ErrAmqpURICannotBeEmpty error = errors.New("amqpURI cannot be empty")

	ErrExchangeNameCannotBeEmpty error = errors.New("exchangeName cannot be empty")

	ErrExchangeTypeCannotBeEmpty error = errors.New("exchangeType cannot be empty")
)
