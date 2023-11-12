package main

import (
	"github.com/sashabaranov/go-openai"

	"GoGPT/router"
)

type Messages struct {
	msg []openai.ChatCompletionMessage
}

func NewMessages() *Messages {
	return &Messages{
		msg: make([]openai.ChatCompletionMessage, 0),
	}
}

func (m *Messages) AddChatMessageRoleUserMsg(content string) {
	m.msg = append(m.msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}

func main() {
	engine := router.Engine()
	_ = engine.Run(":80")

	//config := conf.OpenAIConfig
	//clientConfig := openai.DefaultConfig(config.OpenAI.AuthToken)
	//clientConfig.BaseURL = config.OpenAI.BaseUrl
	//client := openai.NewClientWithConfig(clientConfig)
	//messages := NewMessages()
	//for {
	//	var content string
	//	fmt.Scanln(&content)
	//	if content == "exit" {
	//		break
	//	}
	//	fmt.Println("You: " + content)
	//	messages.AddChatMessageRoleUserMsg(content)
	//	resp, err := client.CreateChatCompletion(
	//		context.Background(),
	//		openai.ChatCompletionRequest{
	//			Model:    openai.GPT3Dot5Turbo,
	//			Messages: messages.msg,
	//		})
	//	if err != nil {
	//		fmt.Printf("ChatCompletion error: %v\n", err)
	//		return
	//	}
	//	fmt.Println(resp.Choices[0].Message.Content)
	//}
}
