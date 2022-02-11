package entity

type ReportForPushToTalk struct {
	ChatName        string `json:"chatroom_name"`
	SenderName      string `json:"sender"`
	RecordingLength string `json:"recording_length"`
}
