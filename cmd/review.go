package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review [file path]",
	Short: "Review a code file using the local LLM",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ファイルパスからコードを読み込む
		filePath := args[0]
		code, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read the file: %v", err)
		}

		// LLMサーバーのエンドポイントを設定
		apiURL := os.Getenv("LLM_API_URL")
		if apiURL == "" {
			log.Fatal("LLM_API_URL is not set")
		}

		// クライアント設定
		config := openai.DefaultConfig("")
		config.BaseURL = apiURL
		client := openai.NewClientWithConfig(config)

		// プロンプトの作成
		prompt := fmt.Sprintf(`You are a code reviewer. Please review the following code and provide feedback on:
- Potential bugs
- Code style issues
- Suggestions for improvement.

Code:
%s`, string(code))

		// LLMに送信
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: "llama3.1-8b-instruct", // ローカルモデルに合わせる
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a code reviewer."},
				{Role: "user", Content: prompt},
			},
		})
		if err != nil {
			log.Fatalf("API call failed: %v", err)
		}

		// レスポンスを表示
		fmt.Println("Code Review Result:")
		fmt.Println(resp.Choices[0].Message.Content)
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}

