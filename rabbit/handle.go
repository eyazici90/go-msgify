package rabbit

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type Handle func(*Message)

func handle(c *context, deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		msg := &Message{}
		_ = json.Unmarshal(d.Body, msg)
		c.handler(msg)
		d.Ack(false)
	}
	done <- nil
}
