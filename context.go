package msgify

import (
	"github.com/streadway/amqp"
)

type context struct {
	conn          *amqp.Connection
	channel       *amqp.Channel
	done          chan error
	amqpURI       string
	exchangeName  string
	exchangeType  string
	queueName     string
	bindingKey    string
	prefetchCount int
	prefetchSize  int
	handler       Handle
}

func NewContext(amqpURI string) RabbitMqContext {
	return &context{
		conn:          nil,
		channel:       nil,
		amqpURI:       amqpURI,
		prefetchCount: DefaultPrefetchCount,
		prefetchSize:  DefaultPrefetchSize,
	}
}

func copy(ctx *context, mutator func(*context)) RabbitMqContext {
	newCtx := &context{
		conn:          ctx.conn,
		channel:       ctx.channel,
		amqpURI:       ctx.amqpURI,
		bindingKey:    ctx.bindingKey,
		exchangeName:  ctx.exchangeName,
		exchangeType:  ctx.exchangeType,
		queueName:     ctx.queueName,
		done:          ctx.done,
		handler:       ctx.handler,
		prefetchCount: ctx.prefetchCount,
		prefetchSize:  ctx.prefetchSize,
	}
	mutator(newCtx)
	return newCtx
}

func WithContext(rCtx RabbitMqContext) RabbitMqContext {
	ctx := rCtx.(*context)
	return copy(ctx, func(*context) {})
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

func (c *context) WithPrefetch(count, size int) RabbitMqContext {
	return copy(c, func(ctx *context) {
		ctx.prefetchCount = count
		ctx.prefetchSize = size
	})
}
