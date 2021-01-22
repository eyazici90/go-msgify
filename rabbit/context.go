package rabbit

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

func NewContext(amqpURI string) MqContext {
	return &context{
		conn:          nil,
		channel:       nil,
		amqpURI:       amqpURI,
		prefetchCount: DefaultPrefetchCount,
		prefetchSize:  DefaultPrefetchSize,
	}
}

func copy(ctx *context, mutator func(*context)) MqContext {
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

func WithContext(rCtx MqContext) MqContext {
	return copy(rCtx.(*context), func(*context) {})
}

func (c *context) WithExchange(name, exchangeType string) MqContext {
	return copy(c, func(ctx *context) {
		ctx.exchangeName = name
		ctx.exchangeType = exchangeType
	})
}

func (c *context) WithQueue(queueName, key string) MqContext {
	return copy(c, func(ctx *context) {
		ctx.queueName = queueName
		ctx.bindingKey = key
	})
}

func (c *context) WithHandle(handle Handle) MqContext {
	return copy(c, func(ctx *context) {
		ctx.handler = handle
	})
}

func (c *context) WithPrefetch(count, size int) MqContext {
	return copy(c, func(ctx *context) {
		ctx.prefetchCount = count
		ctx.prefetchSize = size
	})
}
