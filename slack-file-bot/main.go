package main 

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("key.env")
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
	}

	//set slack-bot-token and channel id
	// os.Setenv("SLACK_BOT_TOKEN","xoxb-5779087348563-5873957135620-QH5iGTh5WsplFxHhrBuaajOC")
	// os.Setenv("CHANNEL_ID","C05NZM8V2SY")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	
	fileArr := []string{"filetoupload.png"}  //need to pass file here


	// this loop is to add multiple files 
	for i := 0; i < len(fileArr); i++ {
		// creating a params for files with channel using external package slack-go/slack
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		// calling UploadFile function from slack api
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n",file.Name, file.URL)
	} 
}