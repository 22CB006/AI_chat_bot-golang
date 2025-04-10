package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Edw590/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Event:")
		fmt.Println("Timestamp:", event.Timestamp)
		fmt.Println("Command:", event.Command)
		fmt.Println("Parameters:", event.Parameters)
		fmt.Println("Event:", event.Event)
		fmt.Println()
	}
}

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file. Make sure it exists.")
	}

	// Init Slack bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Init Wit.ai client
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	// Init Wolfram client
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	// Log command events
	go printCommandEvents(bot.CommandEvents())

	// Bot command definition
	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "Ask a question to Wolfram Alpha",
		Example:     "query for bot - what is the capital of France?",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")

			// Parse with Wit.ai
			msg, err := client.Parse(&witai.MessageRequest{Query: query})
			if err != nil {
				response.Reply("Wit.ai parsing failed.")
				return
			}

			// Extract Wolfram query entity
			data, _ := json.MarshalIndent(msg, "", "    ")
			jsonString := string(data)
			value := gjson.Get(jsonString, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()

			if answer == "" {
				answer = query // fallback
			}

			// Call Wolfram Alpha
			params := url.Values{}
			params.Set("units", "metric")

			queryResult, err := wolframClient.GetQueryResult(answer, params)
			if err != nil {
				fmt.Println("Error from Wolfram:", err)
				response.Reply("Sorry, there was an error retrieving data from Wolfram.")
				return
			}

			if queryResult.QueryResult.Success {
				for _, pod := range queryResult.QueryResult.Pods {
					if pod.Primary || pod.Title == "Result" || pod.Title == "Definition" {
						if len(pod.SubPods) > 0 && pod.SubPods[0].Plaintext != "" {
							response.Reply(pod.SubPods[0].Plaintext)
							return
						}
					}
				}
				response.Reply("I tried, but couldn't find a specific answer.")
			} else {
				response.Reply("Wolfram couldn't find an answer to that question.")
			}
		},
	})

	// Start bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
