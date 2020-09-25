package msgify

import "errors"

var (
	NoChannelFound error = errors.New("No channel found. Please connect to exchange first!")

	AmqpURICannotBeEmpty error = errors.New("amqpURI cannot be empty")

	ExchangeNameCannotBeEmpty error = errors.New("exchangeName cannot be empty")

	ExchangeTypeCannotBeEmpty error = errors.New("exchangeType cannot be empty")
)
