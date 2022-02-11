package repository

import (
	uuid "github.com/satori/go.uuid"

	"github.com/pisitj/app-push-to-talk/2-code/backend/entity"
)

func GetPushToTalk(chat_id uuid.UUID, message_id uuid.UUID) []byte {
}

func CreatePushToTalk(chat_id uuid.UUID) entity.ChatMessage {
}

func GetReportForPushToTalk() []entity.ReportForPushToTalk {
}
