package bot

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"math/rand"
	"time"
)

func answerCallbackMessage(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery) {
	if api, err := bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
		CallbackQueryID: update.ID,
		ShowAlert:       false,
	}); err != nil {
		log.Printf("Error answerCallbackMessage %v\n", api.ErrorCode)
	}
}

func deleteStartMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	del := tgbotapi.DeleteMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
	}

	if _, err := bot.Send(del); err != nil {
		log.Printf("error delete message, userID %v\n", update.Message.Chat.ID)
	}
}

func deleteMessageID(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery, messID int) {

	del := tgbotapi.DeleteMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: messID,
	}

	if _, err := bot.Send(del); err != nil {
		log.Printf("error delete message, userID %v\n", update.Message.Chat.ID)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
