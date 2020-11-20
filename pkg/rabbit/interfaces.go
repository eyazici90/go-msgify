package rabbit

type (
	ContextBuilder interface {
		WithExchange(name, exchangeType string) MqContext
		WithQueue(queueName, key string) MqContext
		WithHandle(handle Handle) MqContext
	}
	Publisher interface {
		Publish(message Message) error
	}
	Consumer interface {
		StartConsuming(consumerTag string) error
		StartConsumingBy(consumerTag string, handle Handle) error
	}
	MqContext interface {
		ContextBuilder
		Connector
		Publisher
		Consumer
	}
)
