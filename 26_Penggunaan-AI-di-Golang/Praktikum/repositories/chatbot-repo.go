package repositories

import (
	"chatbot/dto"

	"context"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo // see another model options: https://platform.openai.com/docs/dto
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func GetLaptopRecommendation(laptop dto.LaptopRequest) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load env file. Err: %s", err)
	}
	content := 
		"Give me a laptop recommendation with a budget of Rp. " +
		strconv.Itoa(laptop.Budget) + " and the purpose is " + laptop.Purpose
		// " Please give the answer with the format: 1. <laptop name>: <Reason>\n2. <laptop name>: <Reason>\n and so on with no opening message"

	ctx := context.Background()
	client := openai.NewClient(os.Getenv("KEY"))
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a friendly chatbot.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	model := openai.GPT3Dot5Turbo
	resp, err := getCompletionFromMessages(ctx, client, messages, model)
	result := resp.Choices[0].Message.Content
	return result, err
}
