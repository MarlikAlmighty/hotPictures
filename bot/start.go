package bot

import (
	"github.com/MarlikAlmighty/picbot/models"
	"github.com/muesli/cache2go"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"time"
)

func start(bot *tgbotapi.BotAPI, update *tgbotapi.Update, cache *cache2go.CacheTable) {

	deleteStartMessage(bot, update)

	title := "What are we going to watch?"
	boobs := "Boobs"
	ass := "Ass"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "<b>"+title+"</b>")
	msg.ParseMode = "HTML"

	rows := make([][]tgbotapi.InlineKeyboardButton, 0, 0)
	row := tgbotapi.NewInlineKeyboardRow()
	row = append(row, tgbotapi.NewInlineKeyboardButtonData(boobs, "Boobs"))
	row = append(row, tgbotapi.NewInlineKeyboardButtonData(ass, "Ass"))
	rows = append(rows, row)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)

	msg.ReplyMarkup = &keyboard

	mess, err := bot.Send(msg)
	if err != nil {
		log.Printf("error send menu to userID %v", update.Message.Chat.ID)
	}

	t := time.Now().Unix()

	m := models.UserCache{
		MessageID: mess.MessageID,
		FloodTime: t,
	}

	cache.Add(update.Message.Chat.ID, 10*time.Minute, &m)
}
