package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".env読み込みエラー: %v", err)
	}
	fmt.Println(".envを読み込みました。")
}

func main() {
	loadEnv()
	var (
		Token = "Bot " + os.Getenv("TOKEN1")
		//BotName = "<@" + os.Getenv("APPLICATION_ID") + ">"
		stopBot = make(chan bool)
		//vcsession         *discordgo.VoiceConnection
		//HelloWorld        = "!helloworld"
		//ChannelVoiceJoin  = "!vcjoin"
		//ChannelVoiceLeave = "!vcleave"
	)

	//fmt.Println(BotName)

	//Discordのセッションを作成
	discord, err := discordgo.New(Token)
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	fmt.Println("Listening...")
	<-stopBot //プログラムが終了しないようロック
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientID := os.Getenv("APPLICATION_ID1")
	u := m.Author
	fmt.Printf("ChannelID: %s, Username: %s(ID: %s) > Content: %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientID {
		sendMessage(s, m.ChannelID, u.Mention()+"なんか喋った!")
		sendReply(s, m.ChannelID, "test", m.Reference())
		outputMessages(s, m)
		newMessage(s, m)
	}
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func sendReply(s *discordgo.Session, channelID string, msg string, reference *discordgo.MessageReference) {
	_, err := s.ChannelMessageSendReply(channelID, msg, reference)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func outputMessages(s *discordgo.Session, m *discordgo.MessageCreate) {

	const limit = 100 // 上限の数を指定

	var (
		beforeID string
		messages []*discordgo.Message
	)

	log.Println("start")

	for {
		c, err := s.ChannelMessages(
			m.ChannelID, // channelID
			limit,       // limit
			beforeID,    // beforeID
			"",          // afterID
			"",          // aroundID
		)
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, c...)

		// limitで指定した件数と一致しない（それ以下）は終了
		if len(c) != limit {
			break
		}

		// 上限まで取得した場合は未取得のものがある可能性が残っているため、
		// 取得した最後のメッセージIDをbeforeIDを設定
		beforeID = c[len(c)-1].ID
	}

	// メッセージの一覧を出力
	for _, msg := range messages {
		fmt.Println(msg.Content)
	}

	log.Println("end")
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "weather"):
		send, err := discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
		log.Println(send.Content)
		if err != nil {
			return
		}
	case strings.Contains(message.Content, "bot"):
		send, err := discord.ChannelMessageSend(message.ChannelID, "Hi there!")
		log.Println(send.Content)
		if err != nil {
			return
		}
	}
}
