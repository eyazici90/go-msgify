package msgify

import (
	"github.com/streadway/amqp"
)

type Connector interface {
	Connect() (RabbitMqContext, error)
	UseConnection(connection *amqp.Connection)
}

func (c *context) Connect() (RabbitMqContext, error) {
	var err error

	if err = verify(c); err != nil {
		return c, err
	}

	if c.conn, err = amqp.Dial(c.amqpURI); err != nil {
		return c, err
	}

	err = returnOnErr(c.openChannel, c.exchangeDeclare)

	return c, err
}

func verify(c *context) error {
	if c.amqpURI == "" {
		return AmqpURICannotBeEmpty
	}
	if c.exchangeName == "" {
		return ExchangeNameCannotBeEmpty
	}
	if c.exchangeType == "" {
		return ExchangeTypeCannotBeEmpty
	}
	return nil
}

func (c *context) UseConnection(connection *amqp.Connection) {
	c.conn = connection
}

func (c *context) openChannel() error {
	var err error
	c.channel, err = c.conn.Channel()

	c.channel.Qos(c.prefetchCount, c.prefetchSize, false)

	return err
}

func (c *context) exchangeDeclare() error {

	if err := c.channel.ExchangeDeclare(
		c.exchangeName, // name of the exchange
		c.exchangeType, // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return err
	}

	return nil
}

func (c *context) bindQueue() error {
	queue, err := c.channel.QueueDeclare(
		c.queueName, // name of the queue
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	if err = c.channel.QueueBind(
		queue.Name,     // name of the queue
		c.bindingKey,   // bindingKey
		c.exchangeName, // sourceExchange
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return err
	}
	return nil
}
