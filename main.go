package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userStates = make(map[int64]string)
var userData = make(map[int64]map[string]string)

func main() {
	bot, err := tgbotapi.NewBotAPI("8181852781:AAGjz8EaeNceExfMWvnEzXtHckCj2U98RA4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			userStates[chatID] = "waiting_for_key"
			userData[chatID] = make(map[string]string)

			msg := tgbotapi.NewMessage(chatID,
				"Отправь мне ключ, который хочешь использовать")
			bot.Send(msg)
			continue
		}

		// Обработка в зависимости от состояния
		switch userStates[chatID] {
		case "waiting_for_key":
			userData[chatID]["key"] = update.Message.Text
			userStates[chatID] = "waiting_for_number"

			msg := tgbotapi.NewMessage(chatID,
				"Отправь мне число, которое хочешь зашифровать")
			bot.Send(msg)

		case "waiting_for_number":
			userData[chatID]["number"] = update.Message.Text
			userStates[chatID] = "waiting_for_module"

			msg := tgbotapi.NewMessage(chatID,
				"Отправь мне модуль, по которому собираешься шифровать")
			bot.Send(msg)

		case "waiting_for_module":
			userData[chatID]["module"] = update.Message.Text
			userStates[chatID] = "main_menu" // ← Новое состояние для меню

			msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Шифровать"),
					tgbotapi.NewKeyboardButton("Создать секретный ключ"),
					tgbotapi.NewKeyboardButton("Завершение"),
				),
			)
			bot.Send(msg)

		case "main_menu": // ← Новый case для обработки кнопок меню
			switch text {
			case "Шифровать":
				// Обработка шифрования
				key, _ := strconv.Atoi(userData[chatID]["key"])
				module, _ := strconv.Atoi(userData[chatID]["module"])
				number, _ := strconv.Atoi(userData[chatID]["number"])

				encryptedMessage := encryption(key, number, module)
				msg := tgbotapi.NewMessage(chatID, "Твоё зашифрованное число: "+strconv.Itoa(encryptedMessage))
				bot.Send(msg)

			case "Создать секретный ключ":
				// Создание секретного ключа
				key, _ := strconv.Atoi(userData[chatID]["key"])
				module, _ := strconv.Atoi(userData[chatID]["module"])

				secretKey := gausMethod(key, fiN(module))
				msg := tgbotapi.NewMessage(chatID, "Твой секретный ключ: "+strconv.Itoa(secretKey))
				bot.Send(msg)

			case "Завершение":
				cleanupEverything(chatID, bot)

				// fmt.Print("Секретный ключ: ")
				// fmt.Print(gausMethod(key, fiN(module)))
				//encryption(key, number, module)
			}
		}
	}
}

func cleanupEverything(chatID int64, bot *tgbotapi.BotAPI) {
	// Очищаем данные и состояния
	delete(userData, chatID)
	delete(userStates, chatID)

	// Убираем клавиатуру
	msg := tgbotapi.NewMessage(chatID, "✅ Все данные очищены!\nДля начала введите /start")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)

	log.Printf("Очищены данные для chatID: %d", chatID)
}
