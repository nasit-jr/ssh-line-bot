package main

import (
	"fmt"
	"net/http"
	"log"
	// "os"

	"github.com/line/line-bot-sdk-go/linebot"
	// "github.com/spf13/viper"
)

func main() {
	// viper.AutomaticEnv()
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.SetDefault("port", "8080")
	// viper.SetConfigType("json")
	// viper.SetConfigName("conf." + os.Getenv("ENV"))
	// viper.AddConfigPath("./conf")



	port := "6500"
	lineSecret := "eb8850e430abea73987642c31811ab23"
	lineToken := "u1bbarCIdvlN/QRq6Zt+vjPkcm9COq9rwF6YOrA3elRde85TaOBiBI+0n2669Vbv1j8OSU+q33hn9sTe5MhSwg9FA8TUJSVZCtO8EvSMMwqIlfdy6VElK8PN/fHIlR2VhT6r4y3+3lQq1OfkfvzvEgdB04t89/1O/w1cDnyilFU="
	
	bot, err := linebot.New(lineSecret, lineToken)
	if err != nil {
		fmt.Println("cannot initiate line client: ", err)
		return
	}



	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			fmt.Println("cannot parse request:", err)
			return
		}
		for _, event := range events {
			fmt.Println(event.Message)
			if event.Type == linebot.EventTypeMessage {
				// message := event.Message
				// str := fmt.Sprintf("%v", message)
				// fmt.Println(str)
				// fmt.Println(message)
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	fmt.Println("Port: ", port)
	http.ListenAndServe(":" + port, nil)
}
