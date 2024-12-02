package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Send a message to the local LLM server",
	Run: func(cmd *cobra.Command, args []string) {
		// ローカルサーバーのエンドポイントを指定
		apiURL := os.Getenv("LLM_API_URL") // 例: http://localhost:8000/v1/chat/completions
		if apiURL == "" {
			log.Fatal("LLM_API_URL is not set")
		}

		// APIキー（ローカルサーバーでは不要な場合も）
		apiKey := os.Getenv("OPENAI_API_KEY")

		// クライアント設定を作成
		config := openai.DefaultConfig(apiKey)
		config.BaseURL = apiURL // ローカルサーバーのURLを設定

		// クライアントの作成
		client := openai.NewClientWithConfig(config)

		// デフォルトメッセージ
		message := "Hello, Local LLM!"
		if len(args) > 0 {
			message = args[0]
		}

		// API呼び出し
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: "llama3.1-8b-instruct", // ローカルモデル名。OpenAI互換で適切にマッピングされる必要あり。
			Messages: []openai.ChatCompletionMessage{
				{Role: "user", Content: message},
			},
		})
		if err != nil {
			log.Fatalf("API call failed: %v", err)
		}

		// レスポンスを表示
		fmt.Println("Local LLM:", resp.Choices[0].Message.Content)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

