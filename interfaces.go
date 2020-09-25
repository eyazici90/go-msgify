package msgify

type (
	ContextBuilder interface {
		WithExchange(name, exchangeType string) RabbitMqContext
		WithQueue(queueName, key string) RabbitMqContext
		WithHandle(handle Handle) RabbitMqContext
	}
	Publisher interface {
		Publish(message Message) error
	}
	Consumer interface {
		StartConsuming(consumerTag string) error
		StartConsumingBy(consumerTag string, handle Handle) error
	}
	RabbitMqContext interface {
		ContextBuilder
		Connector
		Publisher
		Consumer
	}
)
