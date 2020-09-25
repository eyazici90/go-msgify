package msgify

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Message interface {
	GetId() string
	GetMessageType() string
	GetBody() string
	GetContentType() string
	GetCreated() time.Time
}

type message struct {
	Id          string
	MessageType string
	Body        string
	ContentType string
	Created     time.Time
}

func NewMessage(i interface{}) (Message, error) {
	j, err := toJsonDefault(i)
	if err != nil {
		return nil, err
	}

	m := message{
		Id:          uuid.New().String(),
		MessageType: reflect.ValueOf(i).Type().String(),
		Body:        string(j),
		ContentType: defaultContentType,
		Created:     time.Now(),
	}

	return &m, nil
}

func (m *message) GetId() string { return m.Id }

func (m *message) GetMessageType() string { return m.MessageType }

func (m *message) GetBody() string { return m.Body }

func (m *message) GetContentType() string { return m.ContentType }

func (m *message) GetCreated() time.Time { return m.Created }
