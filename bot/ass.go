package bot

import (
	"fmt"
	"github.com/MarlikAlmighty/picbot/models"
	"github.com/muesli/cache2go"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"time"
)

func ass(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery, cache *cache2go.CacheTable) {

	var (
		asslink string
		msg tgbotapi.EditMessageTextConfig
	)

	res, err := cache.Value(update.Message.Chat.ID)
	if err == nil {

		for {

			tmplink := "http://media.obutts.ru/butts_preview/" + fmt.Sprintf("%05d", random(7, 7045)) + ".jpg"

			resp, err := http.Get(tmplink)
			if err != nil {
				continue
			}

			if resp.StatusCode == 200 {
				asslink = tmplink
				break
			}
		}

		title := "What are we going to watch?"
		boobs := "Boobs"
		ass := "Ass"

		body := fmt.Sprintf("<b>%s</b>\n <a href='%s'>&#8203;</a>", title, asslink)

		msg = tgbotapi.NewEditMessageText(update.Message.Chat.ID, res.Data().(*models.UserCache).MessageID, body)
		msg.ParseMode = "HTML"

		rows := make([][]tgbotapi.InlineKeyboardButton, 0, 0)
		row := tgbotapi.NewInlineKeyboardRow()
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(boobs, "Boobs"))
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(ass, "Ass"))
		rows = append(rows, row)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)

		msg.ReplyMarkup = &keyboard

		millis := time.Now().Round(time.Millisecond).UnixNano() / 1e6
		floodTime := res.Data().(*models.UserCache).FloodTime
		dif := millis - floodTime
		if dif < 300 {
			time.Sleep(time.Duration(dif) * time.Millisecond)
		}

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

	} else {

		// TODO Save to database, will solve the problem.

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
}
