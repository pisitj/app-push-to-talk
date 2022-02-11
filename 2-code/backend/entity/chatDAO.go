package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Chat struct {
	chat_id   uuid.UUID
	chat_name string
	member    []uuid.UUID
}

type ChatMessage struct {
	chat_id      uuid.UUID
	message_id   uuid.UUID
	sender       uuid.UUID
	message_type string
	message_body string
	created_at   string
}

type ChatMessagePushToTalk struct {
	message_id     uuid.UUID
	message_url    string
	message_length int
}
