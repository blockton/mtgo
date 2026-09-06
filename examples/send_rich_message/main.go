// Command send_rich_message demonstrates Bot API 10.1 rich messages: sending
// HTML/Markdown rich content, streaming a draft chunk by chunk, and inline
// query results that carry a rich message.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mtgo-labs/mtgo/telegram"
	"github.com/mtgo-labs/mtgo/telegram/types"
)

func main() {
	apiID := mustEnv("API_ID")
	apiHash := mustEnv("API_HASH")
	botToken := mustEnv("BOT_TOKEN")
	chatID, _ := strconv.ParseInt(mustEnv("CHAT_ID"), 10, 64)

	client, err := telegram.NewClient(mustAtoi(apiID), apiHash, &telegram.Config{
		BotToken:    botToken,
		SessionName: "rich_message_bot",
		SavePeers:   true,
	})
	if err != nil {
		log.Fatalf("create client: %v", err)
	}
	if err := client.Start(); err != nil {
		log.Fatalf("start: %v", err)
	}
	defer client.Stop()

	ctx := context.Background()

	// 1. Send a rich message from Markdown markup.
	if _, err := client.SendRichMessageMarkdown(ctx, chatID,
		"# Report\n**bold** and *italic* with `code`"); err != nil {
		log.Fatalf("send rich message: %v", err)
	}

	// 2. Stream a rich message draft the way an AI answer would be streamed.
	draft, err := client.StartRichMessageDraft(ctx, chatID, &telegram.DraftOpts{CanStop: true})
	if err != nil {
		log.Fatalf("start draft: %v", err)
	}
	chunks := []string{"# Streaming", "# Streaming\n\nFirst ", "# Streaming\n\nFirst **chunk** then the rest."}
	for _, c := range chunks {
		if err := draft.SendMarkdown(ctx, c); err != nil {
			log.Fatalf("stream chunk: %v", err)
		}
		time.Sleep(300 * time.Millisecond)
	}
	if err := draft.Stop(ctx); err != nil {
		log.Fatalf("stop draft: %v", err)
	}

	// 3. Answer inline queries with a rich message result.
	client.AddHandler(telegram.NewInlineQueryHandler(func(c *telegram.Context) error {
		return c.AnswerInlineResult(&types.InlineRich{
			ID:          "rich1",
			Title:       c.InlineQuery.Query,
			Description: "Rich message result",
			RichMessage: telegram.RichMessageMarkdown("# "+c.InlineQuery.Query+"\nGenerated content.", false),
		})
	}))

	fmt.Println("rich message demo running; press Ctrl+C to stop")
	client.Idle()
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("environment variable %s is required", key)
	}
	return v
}

func mustAtoi(s string) int32 {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatalf("invalid integer %q: %v", s, err)
	}
	return int32(n)
}
