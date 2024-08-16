package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func Chat(modelname string, msg string) (string, error) {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
	)

	ctx := context.Background()
	fmt.Println("----- standard request -----")
	req := model.ChatCompletionRequest{
		Model: modelname,
		Messages: []*model.ChatCompletionMessage{
			//{
			//	Role: model.ChatMessageRoleSystem,
			//	Content: &model.ChatCompletionMessageContent{
			//		StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
			//	},
			//},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(msg),
				},
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("standard chat error: %v\n", err)
		return "", err
	}
	fmt.Println(*resp.Choices[0].Message.Content.StringValue)
	return *resp.Choices[0].Message.Content.StringValue, nil
}
