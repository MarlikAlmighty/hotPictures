package bot

import (
	"context"
	. "github.com/MarlikAlmighty/picbot/models"
	"github.com/muesli/cache2go"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

func Run(cfg *Config) error {

	var (
		bot *tgbotapi.BotAPI
		err error
	)

	if bot, err = tgbotapi.NewBotAPI(cfg.BotToken); err != nil {
		return err
	}

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(cfg.WebHook)); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache := cache2go.Cache("Cache")

	go loop(ctx, bot, cache)

	log.Printf("Let's go, bot serving on: %s\n", cfg.Host+":"+cfg.Port)

	return http.ListenAndServe(cfg.Host+":"+cfg.Port, nil)
}

func loop(ctx context.Context, bot *tgbotapi.BotAPI, cache *cache2go.CacheTable) {

	updates := bot.ListenForWebhook("/")

	for {

		select {

		case <-ctx.Done():

			return

		case update := <-updates:

			if update.Message != nil {

				switch update.Message.Text {

				default:

					go start(bot, &update, cache)

				}

			} else if update.CallbackQuery != nil {

				answerCallbackMessage(bot, update.CallbackQuery)

				switch update.CallbackQuery.Data {

				case "Boobs":

					go boobs(bot, update.CallbackQuery, cache)

				case "Ass":

					go ass(bot, update.CallbackQuery, cache)
				}
			}
		}
	}
}
