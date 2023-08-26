package main

import (
	"context"

	"log"
	"os"
	"os/signal"

	"kotbot/DbTools"
	"kotbot/handlers"

	"github.com/BurntSushi/toml"
	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var c DbTools.Config
	if _, err := toml.DecodeFile("cfg/local.toml", &c); err != nil {
		panic(err)
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefHandler),
		bot.WithMessageTextHandler("мяу", bot.MatchTypePrefix, handlers.MeowHandler),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, handlers.StartHandler),
		bot.WithMessageTextHandler("отправь:", bot.MatchTypePrefix, handlers.GetPhotoHandler),
	}

	b, err := bot.New(c.Bot.Token, opts...)
	if err != nil {
		log.Panic(err)
	}

	b.Start(ctx)
}
