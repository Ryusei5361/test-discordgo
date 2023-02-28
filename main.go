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

const (
	TOKEN          = "TOKEN1"
	APPLICATION_ID = "APPLICATION_ID1"
)

type stationInfo struct {
	info  string
	count int
}

func main() {
	loadEnv()
	var (
		Token = "Bot " + os.Getenv(TOKEN)
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

	discord.AddHandler(onMessageUpdate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {

		}
	}(discord)

	fmt.Println("Listening...")
	<-stopBot //プログラムが終了しないようロック
	return
}

func onMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	fmt.Println("<start>")
	const limit = 1 // 上限の数を指定

	c, err := s.ChannelMessages(
		m.ChannelID, // channelID
		limit,       // limit
		"",          // beforeID
		"",          // afterID
		"",          // aroundID
	)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(c[0].Content, "\n")
	var contentInfo []stationInfo
	var content []string

	for i := range contents {
		if i%3 != 0 {
			content = append(content, contents[i])
		}
	}

	for _, i := range content {
		//fmt.Printf("i: %v\n", i)
		for _, j := range contentInfo {
			//fmt.Printf("j: %v\n", j)

			if contains(i) {
				//content := strings.Fields(content[i])
				//station := strings.Join(content[0:3], " ")
				//price, _ := strconv.Atoi(strings.Join(content[3:4], " "))
				contentInfo = append(contentInfo, stationInfo{info: i})
				break
			}
		}
	}
	fmt.Println(contentInfo)
	fmt.Println("<end>")
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientID := os.Getenv(APPLICATION_ID)
	u := m.Author
	fmt.Printf("ChannelID: %s, Username: %s(ID: %s) > Content: %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientID {
		//sendMessage(s, m.ChannelID, u.Mention()+"なんか喋った!")
		//sendReply(s, m.ChannelID, "test", m.Reference())
		outputMessages(s, m)
		//newMessage(s, m)
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

	const limit = 2 // 上限の数を指定

	var (
		beforeID string
		//messages []*discordgo.Message
	)

	log.Println("start")

	//for {
	//	c, err := s.ChannelMessages(
	//		m.ChannelID, // channelID
	//		limit,       // limit
	//		beforeID,    // beforeID
	//		"",          // afterID
	//		"",          // aroundID
	//	)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	messages = append(messages, c...)
	//
	//	// limitで指定した件数と一致しない（それ以下）は終了
	//	if len(c) != limit {
	//		break
	//	}
	//
	//	// 上限まで取得した場合は未取得のものがある可能性が残っているため、
	//	// 取得した最後のメッセージIDをbeforeIDを設定
	//	beforeID = c[len(c)-1].ID
	//}
	//fmt.Println(m.Channel)

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

	fmt.Println(c[0])

	//messages = append(messages, c...)
	//
	//// メッセージの一覧を出力
	//for _, msg := range messages {
	//	fmt.Println(msg.Content)
	//}

	log.Println("end")
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch message.Content {
	case "なにしてたの？":
		send, err := discord.ChannelMessageSend(message.ChannelID, "ねてたぁ")
		log.Println(send.Content)
		if err != nil {
			return
		}
	default:
		send, err := discord.ChannelMessageSend(message.ChannelID, "しぬかぁ")
		log.Println(send.Content)
		if err != nil {
			return
		}
	}
	// Respond to messages
	//switch {
	//case strings.Contains(message.Content, "かんやのあほ"):
	//	send, err := discord.ChannelMessageSend(message.ChannelID, "うるせぇ")
	//	log.Println(send.Content)
	//	if err != nil {
	//		return
	//	}
	//case strings.Contains(message.Content, "サッカー好き？"):
	//	send, err := discord.ChannelMessageSend(message.ChannelID, " 大好き！！！")
	//	log.Println(send.Content)
	//	if err != nil {
	//		return
	//	}
	//case strings.Contains(message.Content, "ゆうみちゃん好き？"):
	//	send, err := discord.ChannelMessageSend(message.ChannelID, " 愛してる！！！")
	//	log.Println(send.Content)
	//	if err != nil {
	//		return
	//	}
	//case strings.Contains(message.Content, "かんや"):
	//	send, err := discord.ChannelMessageSend(message.ChannelID, "なに？")
	//	log.Println(send.Content)
	//	if err != nil {
	//		return
	//	}
	//case strings.Contains(message.Content, "かにゃ"):
	//	send, err := discord.ChannelMessageSend(message.ChannelID, "はにゃ？")
	//	log.Println(send.Content)
	//	if err != nil {
	//		return
	//	}
	//}
}
