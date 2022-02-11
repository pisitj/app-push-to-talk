package repository

import (
	uuid "github.com/satori/go.uuid"

	"github.com/pisitj/app-push-to-talk/2-code/backend/entity"
)

func GetChatList(user_id uuid.UUID) []entity.Chat {
}

func CreateNewChat(member []uuid.UUID) entity.Chat {
}

func GetChatHistory(user_id uuid.UUID, chat_id uuid.UUID) []entity.ChatMessage {
}

func CreateChatMessage(sender uuid.UUID, chat_id uuid.UUID, message_type string, message_body string) entity.ChatMessage {
}
