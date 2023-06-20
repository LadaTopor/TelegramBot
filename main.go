package main

import (
	"database/sql"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"time"
)

func main() {
	db, err := PostgresConnection()
	if err != nil {
		log.Fatal("DB connection error")
	}

	bot, err := tgbotapi.NewBotAPI("6031585136:AAGiev9aBT3qIRPJavMToB1SBSJkd9AgknI")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			// Обработка нажатия кнопок встроенного меню
			callback := update.CallbackQuery
			switch callback.Data {
			default:
				weekday, schedule := selectWeekday(callback.Data, db)
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("%s\n%s", weekday, schedule))
				bot.Send(msg)
			}

			// Отправка ответа на нажатие кнопки
			answerCallback := tgbotapi.NewCallback(callback.ID, "")
			bot.AnswerCallbackQuery(answerCallback)
		}
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/start" {
			replyMarkup := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("/allDays"),
					tgbotapi.NewKeyboardButton("/today"),
					tgbotapi.NewKeyboardButton("/selectDay"),
				),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот-расписание, могу выдать расписание на любой день недели!")
			msg.ReplyMarkup = replyMarkup

			_, err := bot.Send(msg)
			if err != nil {
				log.Fatal(err)
			}

		} else if update.Message.Text == "/allDays" {
			rows, err := db.Query(`SELECT weekday, schedule FROM timetable ORDER BY id ASC`)
			if err != nil {
				log.Fatal(err)
			}

			var weekday, schedule string
			for rows.Next() {
				err = rows.Scan(&weekday, &schedule)
				if err != nil {
					log.Fatal(err)
				}
				msgWeekday := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s", weekday, schedule))

				_, err := bot.Send(msgWeekday)
				if err != nil {
					log.Println(err)
				}
			}
		} else if update.Message.Text == "/today" {
			daysOfWeek := map[time.Weekday]string{
				time.Monday:    "ПН",
				time.Tuesday:   "ВТ",
				time.Wednesday: "СР",
				time.Thursday:  "ЧТ",
				time.Friday:    "ПТ",
			}
			today := time.Now().Weekday()

			var weekday, schedule string
			err := db.QueryRow(`SELECT weekday, schedule FROM timetable WHERE weekday = $1`, daysOfWeek[today]).Scan(
				&weekday, &schedule)
			if err != nil {
				log.Fatal(err)
			}
			msgWeekday := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s", weekday, schedule))
			_, err = bot.Send(msgWeekday)
			if err != nil {
				log.Println(err)

			}
		} else if update.Message.Text == "/selectDay" {
			inlineRow := []tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("Понедельник", "monday"),
				tgbotapi.NewInlineKeyboardButtonData("Вторник", "tuesday"),
				tgbotapi.NewInlineKeyboardButtonData("Среда", "wednesday"),
				tgbotapi.NewInlineKeyboardButtonData("Четверг", "thursday"),
				tgbotapi.NewInlineKeyboardButtonData("Пятница", "friday"),
			}
			inlineMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineRow)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите день недели:")
			msg.ReplyMarkup = inlineMarkup
			//log.Println("hello222bich")

			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}

	}
}

func selectWeekday(day string, db *sql.DB) (string, string) {
	daysOfWeek := map[string]string{
		"monday":    "ПН",
		"tuesday":   "ВТ",
		"wednesday": "СР",
		"thursday":  "ЧТ",
		"friday":    "ПТ",
	}
	var weekday, schedule string
	err := db.QueryRow(`SELECT weekday, schedule FROM timetable WHERE weekday = $1`, daysOfWeek[day]).Scan(
		&weekday, &schedule)
	if err != nil {
		log.Fatal(err)
	}
	return weekday, schedule
}
