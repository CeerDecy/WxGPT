package gptclient

import (
	"context"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"

	"WxGPT/conf"
	"WxGPT/internal/gpt/message"
)

type GptClient struct {
	c *openai.Client
}

func DefaultClient() *GptClient {
	config := conf.OpenAIConfig
	clientConfig := openai.DefaultConfig(config.OpenAI.AuthToken)
	clientConfig.BaseURL = config.OpenAI.BaseUrl
	client := openai.NewClientWithConfig(clientConfig)
	return &GptClient{
		c: client,
	}
}

func (g *GptClient) GetStreamResponse(content string) (openai.ChatCompletionStreamResponse, error) {
	msgs := message.NewMessages()
	msgs.AddChatMessageRoleUserMsg(content)
	stream, err := g.c.CreateChatCompletionStream(context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: msgs.Msg,
			Stream:   true,
		})
	if err != nil {
		log.Println(err)
	}
	response, err := stream.Recv()
	//response.Object
	//io.ReadAll()
	return response, err
}

func (g *GptClient) GetResponse(content string) (string, error) {
	msgs := message.NewMessages()
	msgs.AddChatMessageRoleUserMsg(content)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*4500)
	defer cancelFunc()
	resp, err := g.c.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: msgs.Msg,
			//Stream:   true,
		})
	if err != nil {
		return err.Error(), err
	}
	//for i := range recv.Object {
	//	fmt.Println(recv.Object[i])
	//}
	//resp, err := g.c.CreateChatCompletion(
	//	context.Background(),
	//	openai.ChatCompletionRequest{
	//		Model:    openai.GPT3Dot5Turbo,
	//		Messages: msgs.Msg,
	//	})
	return resp.Choices[0].Message.Content, err
	//return "", err
}

//
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
