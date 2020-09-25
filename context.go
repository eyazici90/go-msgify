package msgify

import (
	"github.com/streadway/amqp"
)

type context struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	amqpURI      string
	exchangeName string
	exchangeType string
	queueName    string
	bindingKey   string
	handler      Handle
}

func NewContext(amqpURI string) RabbitMqContext {
	return &context{
		conn:    nil,
		channel: nil,
		amqpURI: amqpURI,
	}
}

func WithContext(rCtx RabbitMqContext) RabbitMqContext {
	ctx := rCtx.(*context)

	c := &context{
		conn:         ctx.conn,
		channel:      ctx.channel,
		amqpURI:      ctx.amqpURI,
		bindingKey:   ctx.bindingKey,
		exchangeName: ctx.exchangeName,
		exchangeType: ctx.exchangeType,
		queueName:    ctx.queueName,
		done:         ctx.done,
		handler:      ctx.handler,
	}
	return c
}

func copy(ctx *context, mutator func(*context)) RabbitMqContext {
	newCtx := &context{
		conn:         ctx.conn,
		channel:      ctx.channel,
		amqpURI:      ctx.amqpURI,
		bindingKey:   ctx.bindingKey,
		exchangeName: ctx.exchangeName,
		exchangeType: ctx.exchangeType,
		queueName:    ctx.queueName,
		done:         ctx.done,
		handler:      ctx.handler,
	}
	mutator(newCtx)
	return newCtx
}

func (c *context) WithExchange(name, exchangeType string) RabbitMqContext {
	return copy(c, func(ctx *context) {
		ctx.exchangeName = name
		ctx.exchangeType = exchangeType
	})
}

func (c *context) WithQueue(queueName, key string) RabbitMqContext {
	return copy(c, func(ctx *context) {
		ctx.queueName = queueName
		ctx.bindingKey = key
	})
}

func (c *context) WithHandle(handle Handle) RabbitMqContext {
	return copy(c, func(ctx *context) {
		ctx.handler = handle
	})
}
