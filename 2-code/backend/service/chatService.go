package service

import (
	uuid "github.com/satori/go.uuid"

	"github.com/pisitj/app-push-to-talk/2-code/backend/entity"
	"github.com/pisitj/app-push-to-talk/2-code/backend/repository"
)

func GetChatList(user_id uuid.UUID) []entity.Chat {
	chatList := repository.GetChatList(user_id)
	return chatList
}

func CreateNewChat(member []uuid.UUID) entity.Chat {
	newChat := repository.CreateNewChat(member)
	return newChat
}

func GetChatHistory(user_id uuid.UUID, chat_id uuid.UUID) []entity.ChatMessage {
	chatHistory := repository.GetChatHistory(user_id, chat_id)
	return chatHistory
}

func CreateChatMessage(sender uuid.UUID, chat_id uuid.UUID, message_type string, message_body string) entity.ChatMessage {
	newChatMessage := repository.CreateChatMessage(sender, chat_id, message_type, message_body)
	return newChatMessage
}
