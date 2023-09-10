package main

import(
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	err := godotenv.Load("key.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// os.Setenv("SLACK_BOT_TOKEN","xoxb-5779087348563-5810572203569-nTTNSZJsY9tWswd0rkbhrDDr")
	// os.Setenv("SLACK_APP_TOKEN","xapp-1-A05PFSHDHQS-5859077653303-1196c9b72feb781181034b3c2ff269e634b0817b2d9c9e57f20d6b57513c3578")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example:    "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year:= request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r:= fmt.Sprintf("age is %d", age)
			response.Reply(r)	
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	
	defer cancel()

	err1 := bot.Listen(ctx)
	if err1 != nil{
		log.Fatal(err1)
	}
}