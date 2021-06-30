package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	USD string = "USD"
	RUB string = "RUB"
)

type bnResp struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

type wallet map[string]float64

var db = map[int64]wallet{}

func exchangeRate(curFrom string, curTo string) (amountTo float64, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sT%s", curFrom, curTo))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var jsonResp bnResp
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		log.Fatal(err)
	}
	if jsonResp.Code != 0 {
		err = errors.New("Неверный символ")
	}
	amountTo = jsonResp.Price
	return
}

func getPrice(symbol string) (price float64, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", symbol))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var jsonResp bnResp

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		log.Fatal(err)
	}

	if jsonResp.Code != 0 {
		err = errors.New("Неверный символ")
	}
	price = jsonResp.Price
	return
}

func main() {
	bot, err := tgbotapi.NewBotAPI("1818286792:AAGI1MSvXv5A10Pp9QRFXh4-DZc8mqS3YnU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	rubTUsd, _ := exchangeRate(USD, RUB)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		command := strings.Split(update.Message.Text, " ")

		switch command[0] {
		case "ADD":
			if len(command) != 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда"))
			}
			amount, err := strconv.ParseFloat(command[2], 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			}
			if _, ok := db[update.Message.Chat.ID]; !ok {
				db[update.Message.Chat.ID] = wallet{}
			}
			db[update.Message.Chat.ID][command[1]] += amount
			balanceText := fmt.Sprintf("%f", db[update.Message.Chat.ID][command[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, balanceText))
		case "SUB":
			if len(command) != 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда"))
			}
			amount, err := strconv.ParseFloat(command[2], 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			}
			if _, ok := db[update.Message.Chat.ID]; !ok {
				continue
			}
			db[update.Message.Chat.ID][command[1]] -= amount
			balanceText := fmt.Sprintf("%f", db[update.Message.Chat.ID][command[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, balanceText))
		case "DEL":
			if len(command) != 2 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда"))
			}
			delete(db[update.Message.Chat.ID], command[1])
		case "SHOW":
			msg := ""
			var sum float64 = 0
			for key, value := range db[update.Message.Chat.ID] {
				price, _ := getPrice(key)
				sum += value * price
				msg += fmt.Sprintf("%s: %f [%.2fUSD/%.2fRUB]\n", key, value, value*price, rubTUsd*value*price)
			}
			msg += fmt.Sprintf("Total: [%.2fUSD/%.2fRUB]\n", sum, sum*rubTUsd)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Команда не найдена"))
		}
	}
}
