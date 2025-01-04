package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Text == "/start" {
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Неделя", "subscription_week"),
					tgbotapi.NewInlineKeyboardButtonData("Месяц", "subscription_month"),
					tgbotapi.NewInlineKeyboardButtonData("Год", "subscription_year"),
				),
			)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите срок подписки:")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		data := update.CallbackQuery.Data
		var response string
		switch data {
		case "subscription_week":
			response = "Вы выбрали подписку на неделю"
		case "subscription_month":
			response = "Вы выбрали подписку на месяц"
		case "subscription_year":
			response = "Вы выбрали подписку на год!"
		default:
			response = "Неизвестный выбор."
		}

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, response)
		bot.Request(callback)

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, response)
		bot.Send(msg)
	}
}
