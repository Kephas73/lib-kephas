package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
)

type WebhookProcessInstance interface {
	MessageCallback(ctx context.Context, b *bot.Bot, update *models.Update)
}

func (t *Telegram) InstallWebhook(processingInstance WebhookProcessInstance) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(processingInstance.MessageCallback),
	}

	b, err := bot.New(t.Config.Bot.Token, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
