package service

import (
	uuid "github.com/satori/go.uuid"

	"github.com/pisitj/app-push-to-talk/2-code/backend/entity"
	"github.com/pisitj/app-push-to-talk/2-code/backend/repository"
)

func GetPushToTalk(chat_id uuid.UUID, message_id uuid.UUID) []byte {
	pushToTalkFile := repository.GetPushToTalk(chat_id, message_id)
	return pushToTalkFile
}

func CreatePushToTalk(chat_id uuid.UUID) entity.ChatMessage {
	chatMessagePushToTalk := repository.CreatePushToTalk(chat_id)
	return chatMessagePushToTalk
}

func GetReportForPushToTalk() []entity.ReportForPushToTalk {
	reportForPushToTalkList := GetReportForPushToTalk()
	return reportForPushToTalkList
}
