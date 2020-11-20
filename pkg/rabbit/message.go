package rabbit

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          string
	MessageType string
	Body        string
	ContentType string
	Created     time.Time
}

func NewMessage(i interface{}) (*Message, error) {
	j, err := toJSONDefault(i)
	if err != nil {
		return nil, err
	}

	m := Message{
		ID:          uuid.New().String(),
		MessageType: reflect.ValueOf(i).Type().String(),
		Body:        string(j),
		ContentType: defaultContentType,
		Created:     time.Now(),
	}

	return &m, nil
}
