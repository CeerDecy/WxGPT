package message

import "github.com/sashabaranov/go-openai"

type Messages struct {
	Msg []openai.ChatCompletionMessage
}

func NewMessages() *Messages {
	return &Messages{
		Msg: make([]openai.ChatCompletionMessage, 0),
	}
}

func (m *Messages) AddChatMessageRoleUserMsg(content string) {
	m.Msg = append(m.Msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}
