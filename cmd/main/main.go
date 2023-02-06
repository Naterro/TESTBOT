package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME
	bot, err := tgbotapi.NewBotAPI("5898000564:AAEH0g-eQN7K1f3oN6J3dKLjkTfPzQTYu2w")
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for v := range updates {
		if v.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(v.Message.Chat.ID, "")
		if !v.Message.IsCommand() { // ignore any non-command Messages
			msg.Text = v.Message.Chat.UserName + ", " + v.Message.Text
			msg.ReplyToMessageID = v.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
		switch v.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	/*
		Umap := tg_api.GetNewMsg()
		for i, v := range Umap {
			switch t := v.(type) {
			case map[string]string:
				fmt.Println("a", t)
			}
			str := fmt.Sprintf("%v", v)
			fmt.Println(i, str)
		}*/
	//tg_api.SendMessage("фывфы", "504416149")

}
