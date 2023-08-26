package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"kotbot/DbTools"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	fmt.Println("DefH ", update.Message.From.ID)
	var text = ""
	if update.Message != nil && update.Message.Photo != nil {
		fmt.Println("DefH ", update.Message.Chat.ID)
		fmt.Println("DefH ", update.Message.Photo[len(update.Message.Photo)-1].FileID)
		words := strings.Split(update.Message.Caption, " ")
		if len(words) == 2 && words[0] == "кот" {
			fmt.Println("added cat")
			var photo = update.Message.Photo[len(update.Message.Photo)-1].FileID
			DbTools.AddCatPhoto(photo, words[1])
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:                update.Message.Chat.ID,
				Text:                  "Фото успешно добавлено!",
				ReplyToMessageID:      update.Message.ID,
				ParseMode:             "",
				Entities:              nil,
				DisableWebPagePreview: false,
			})
			if err != nil {
				panic(err)
			}

		} else {
			text = "Кажется, вы неправильно указали породу. " +
				"Если порода кота состоит более чем из одного слова - " +
				"замените пробелы между словами знаком '_'"
			go b.SendChatAction(ctx, &bot.SendChatActionParams{
				ChatID:          update.Message.Chat.ID,
				MessageThreadID: 0,
				Action:          "typing",
			})
			time.Sleep(2 * time.Second)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:                update.Message.Chat.ID,
				Text:                  text,
				ReplyToMessageID:      update.Message.ID,
				ParseMode:             "",
				Entities:              nil,
				DisableWebPagePreview: false,
			})
			if err != nil {
				panic(err)
			}
		}

	} else {
		text = "МЯЯУУ"
		go b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID:          update.Message.Chat.ID,
			MessageThreadID: 0,
			Action:          "typing",
		})
		time.Sleep(2 * time.Second)
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:                update.Message.Chat.ID,
			Text:                  text,
			ReplyToMessageID:      update.Message.ID,
			ParseMode:             "",
			Entities:              nil,
			DisableWebPagePreview: false,
		})
		if err != nil {
			panic(err)
		}
	}
}

func MeowHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("meowH ", update.Message.From.ID)
	var text = "Нет, блин, гав"

	go b.SendChatAction(ctx, &bot.SendChatActionParams{
		ChatID:          update.Message.Chat.ID,
		MessageThreadID: 0,
		Action:          "typing",
	})
	time.Sleep(2 * time.Second)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:                update.Message.Chat.ID,
		Text:                  text,
		ReplyToMessageID:      update.Message.ID,
		ParseMode:             "",
		Entities:              nil,
		DisableWebPagePreview: false,
	})
	if err != nil {
		panic(err)
	}
}
func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("StartH ", update.Message.From.ID)
	var user_name = update.Message.From.Username
	var text = ""
	if user_name != "" {
		text = "Привет, " + user_name + ", я кошкобот"
	} else {
		text = "Привет, котик, я кошкобот"
	}
	go b.SendChatAction(ctx, &bot.SendChatActionParams{
		ChatID:          update.Message.Chat.ID,
		MessageThreadID: 0,
		Action:          "typing",
	})
	time.Sleep(2 * time.Second)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:                update.Message.Chat.ID,
		Text:                  text,
		ParseMode:             "",
		Entities:              nil,
		DisableWebPagePreview: false,
	})
	if err != nil {
		panic(err)
	}
}
func GetPhotoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("GetPhotoH ", update.Message.From.ID)
	var text = ""
	splitStr := strings.Split(update.Message.Text, ":")
	if len(splitStr) == 2 && len(strings.Split(splitStr[1], " ")) == 1 {
		var photocat = DbTools.GetCatPhoto(splitStr[1])
		fmt.Println("GetPhotoH ", photocat)
		if photocat != "no" {
			_, err := b.SendPhoto(ctx, &bot.SendPhotoParams{
				ChatID:  update.Message.Chat.ID,
				Caption: "Взгляните на этого котика",
				Photo:   &models.InputFileString{Data: DbTools.GetCatPhoto(splitStr[1])},
			})
			if err != nil {
				panic(err)
			}
		} else {
			text = "Такой породы ещё нет в базе"
			go b.SendChatAction(ctx, &bot.SendChatActionParams{
				ChatID:          update.Message.Chat.ID,
				MessageThreadID: 0,
				Action:          "typing",
			})
			time.Sleep(2 * time.Second)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:                update.Message.Chat.ID,
				Text:                  text,
				ReplyToMessageID:      update.Message.ID,
				ParseMode:             "",
				Entities:              nil,
				DisableWebPagePreview: false,
			})
			if err != nil {
				panic(err)
			}
		}
	} else {
		text = "неправильный запрос"
		go b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID:          update.Message.Chat.ID,
			MessageThreadID: 0,
			Action:          "typing",
		})
		time.Sleep(2 * time.Second)
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:                update.Message.Chat.ID,
			Text:                  text,
			ReplyToMessageID:      update.Message.ID,
			ParseMode:             "",
			Entities:              nil,
			DisableWebPagePreview: false,
		})
		if err != nil {
			panic(err)
		}
	}
}
