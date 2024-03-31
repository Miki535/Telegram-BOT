package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	botToken := "6669079500:AAFFkfdH4m1P4Wm5Im-Mc21x2bfOv44gLKs"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Our site").WithText("https://github.com/Miki535/Telegram-BOT"),
				tu.KeyboardButton("Documentation").WithText("To start enter any word in message and tap enter!"),
			),
		)
		message := tu.Message(
			chatId,
			"Hello!",
		).WithReplyMarkup(keyboard)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)
			uu := update.Message.Text
			if uu == "To start enter any word in message and tap enter!" || uu == "https://github.com/Miki535/Telegram-BOT" {
				return
			} else {

				messlen := len(uu)
				mess := fmt.Sprint(messlen)
				fullmess := fmt.Sprintf("Len of youre message..." + mess)

				message := tu.Message(
					chatID,
					fullmess,
				)

				bot.SendMessage(message)

			}
		}
	}, th.AnyMessageWithText())

	bh.Start()
}

//https://telego.pixelbox.dev/docs/introduction/tutorial/
