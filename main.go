package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
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
	TOKEN          = "TOKEN1" // ボット トークンは、API リクエストを承認するために使用される
	APPLICATION_ID = "APPLICATION_ID1"
)

type stationInfo struct {
	station string
	price   string
	count   int
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
	f, err := os.Open("write.txt")
	data := make([]byte, 1024)
	read, err := f.Read(data)
	//fmt.Printf("%s%s", "!!!!!!", string(data[:read]))
	if err != nil {
		return
	}
	for _, r := range string(data[:read]) {
		fmt.Println("!!!!")
		fmt.Println(r)
	}
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
	//var contentInfo []stationInfo
	ca := make(map[string]int)

	for i, content := range contents {
		if i%3 != 0 {
			content := strings.Split(content, ":")
			ca[content[0]+content[1]]++
			//contentInfo = append(contentInfo, stationInfo{station: content[0], price: content[1], count: 0})
		}
	}
	//contentInfo = countStations(contentInfo)

	//fmt.Println(ca)

	f, err := os.Create("write.txt")
	for i, x := range ca {
		data := []byte(i)
		data = append(data, " "...)
		data = append(data, i2s(x)...)
		data = append(data, "\n"...)
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err)
			fmt.Println("fail to write file")
		}
	}

	defer f.Close()
	fmt.Println("<end>")

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientID := os.Getenv(APPLICATION_ID)
	u := m.Author
	fmt.Printf("ChannelID: %s, Username: %s(ID: %s) > Content: %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientID {
		//sendMessage(s, m.ChannelID, u.Mention()+"なんか喋った!")
		//sendReply(s, m.ChannelID, "test", m.Reference())
		//outputMessages(s, m)
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

func countStations(slice []stationInfo) []stationInfo {
	cop := slice
	var ret []stationInfo

	for _, x := range slice {
		count := 0
		for _, y := range cop {
			if x.station == y.station && x.price == y.price {
				count += 1
			}
		}
		if count != 0 {
			ret = append(ret, stationInfo{station: x.station, price: x.price, count: count})
		}
		cop = deleteStations(cop, x)
	}
	return ret
}

func deleteStations(slice []stationInfo, s stationInfo) []stationInfo {
	ret := make([]stationInfo, len(slice))
	i := 0
	for _, x := range slice {
		if s.station != x.station || s.price != x.price {
			ret[i] = x
			i++
		}
	}
	return ret[:i]
}

// String -> Int
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

// Int -> String
func i2s(i int) string {
	return strconv.Itoa(i)
}
